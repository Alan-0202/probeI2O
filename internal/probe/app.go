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
	"time"
)

const (
	TCP_POST_URL = "http://21.64.64.200:8018/v1/hc/network/tcp/portstatus"
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
	client := &http.Client{Timeout: 5 * time.Second, Transport: tr}

	return &appProbe{
		client: client,
		url: TCP_POST_URL,
		header: CONTENT_TYPE,
	}
}

func (ap *appProbe) Post(data interface{}) {
	fmt.Println(data.(model.Source))
	allBytes, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		fmt.Println("parse the data ")
		g.MetricsReqMap.Store(data.(model.Source).Address, "0")
		return
	}

	res, err := ap.client.Post(ap.url, ap.header, bytes.NewBuffer(allBytes))
	if err != nil {
		log.Error(err)
		fmt.Println("tijiao post")
		g.MetricsReqMap.Store(data.(model.Source).Address, "0")
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
		fmt.Println(resBad.Msg)
		fmt.Println(res.StatusCode)
		err := json.Unmarshal(bytes, &resBad)
		if err != nil {
			log.Error(err)
			g.MetricsReqMap.Store(data.(model.Source).Address, "0.5")
			return
		}
		fmt.Println("budegnyu 200")
		g.MetricsReqMap.Store(data.(model.Source).Address, "0")
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
		if ration * 100 < data.(model.Source).Threshold {

			metricsOptsSlice = append(metricsOptsSlice, model.MetricsOpts{
				Ip:     item.Ip,
				Port:   item.Port,
				Meg:    "Ration is not reached !",
				Value:  0,
				Ration: ration * 100,
				Shop:   g.MetricsIpShop[data.(model.Source).Address],
			})
			g.MetricsMap.Store(data.(model.Source).Address, metricsOptsSlice)
		}

		metricsOptsSlice = append(metricsOptsSlice, model.MetricsOpts{
			Ip:     item.Ip,
			Port:   item.Port,
			Meg:    "Ration is reached !",
			Value:  1,
			Ration: ration * 100,
			Shop:   g.MetricsIpShop[data.(model.Source).Address],
		})
		g.MetricsMap.Store(data.(model.Source).Address, metricsOptsSlice)
	}

	return
}
