package probe

import (
	"I2Oprobe/internal/g"
	"I2Oprobe/internal/log"
	"I2Oprobe/internal/model"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	//TCP_POST_URL = "http://xxxxx/v1/hc/network/tcp/portstatus"
	CONTENT_TYPE = "application/json"
)

type appProbe struct {
	client *http.Client
	url    string
	header string
}

func NewAppProbe() *appProbe {
	// init client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout:  *g.OneProbeClientTime , Transport: tr}

	return &appProbe{
		client: client,
		url: *g.ProbeURL,
		header: CONTENT_TYPE,
	}
}

func (ap *appProbe) Post(data interface{}) {
	fmt.Println(data.(model.Source))
	allBytes, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		g.MetricsReqMap.Store(data.(model.Source).Address , "0.5")
		return
	}

	res, err := ap.client.Post(ap.url, ap.header, bytes.NewBuffer(allBytes))
	if err != nil {
		log.Error(err)
		g.MetricsReqMap.Store(data.(model.Source).Address, "0.5")
		return
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		g.MetricsReqMap.Store(data.(model.Source).Address, "0.5")
		return
	}

	if res.StatusCode != 200 {
		var resBad model.ResBad
		err := json.Unmarshal(bytes, &resBad)
		if err != nil {
			log.Error(err)
			g.MetricsReqMap.Store(data.(model.Source).Address, "0.5")
			return
		}
		log.Infof("%v request code != 200.",data.(model.Source).Address)
		g.MetricsReqMap.Store(data.(model.Source).Address, "0")
		return
	}

	var resOk model.RespOk
	err = json.Unmarshal(bytes, &resOk)
	if err != nil {
		log.Error(err)
		g.MetricsReqMap.Store(data.(model.Source).Address, "0.5")
		return
	}
	var metricsOptsSlice []model.MetricsOpts
	for _, item := range resOk.Status {
		ration := item.Ratio
		if ration * 100 < float64(data.(model.Source).Threshold) {
			fmt.Printf("IP: %v, destRation: %v, getRation: %v\n", item.Ip, ration*100, float64(data.(model.Source).Threshold))
			metricsOptsSlice = append(metricsOptsSlice, model.MetricsOpts{
				Address: data.(model.Source).Address,
				Ip:     item.Ip,
				Port:   item.Port,
				Meg:    "Ration is not reached !",
				Value:  -1,
				Ration: ration * 100,
				Shop:   strings.Replace(g.MetricsIpShop[data.(model.Source).Address], "\n", "", -1),
			})
			g.MetricsMap.Store(data.(model.Source).Address, metricsOptsSlice)
			continue
		}

		metricsOptsSlice = append(metricsOptsSlice, model.MetricsOpts{
			Address: data.(model.Source).Address,
			Ip:     item.Ip,
			Port:   item.Port,
			Meg:    "Ration is reached !",
			Value:  1,
			Ration: ration * 100,
			Shop:   strings.Replace(g.MetricsIpShop[data.(model.Source).Address], "\n", "", -1),
		})
		g.MetricsMap.Store(data.(model.Source).Address, metricsOptsSlice)
	}
	return
}
