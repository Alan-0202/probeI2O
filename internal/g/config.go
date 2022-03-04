package g

import (
	"I2Oprobe/internal/model"
	"github.com/alecthomas/kingpin"
	"sync"
)

const version = "1.0.0.0"


// for scrape
var (
	MetricsMap sync.Map
	//MetricsMapCache sync.Map
	MetricsReqMap sync.Map
	//MetricsReqMapCache sync.Map
)
var (
	MetricsIpShop = make(map[string]string)
)

// from file text
var (
	MetricsChan = make(chan model.Source, 10000)
)

var (
	ListenAndPort = kingpin.Flag("listen_addr", "health_check exporter listen addr").Default(":9200").String()
	ConcurNum = kingpin.Flag("concur_num", "goroutine number for post").Default("20").Int()
	//ClientTimeOut = kingpin.Flag("cli_timeout", "request the probeAPI client timeout(default: 3s)").Default("3s").Duration()
	SourcePath = kingpin.Flag("source_path", "get the process object").Default("/home/alan/goworkspace/src/I2Oprobe/probelist.txt").String()
	OneProbeClientTime = kingpin.Flag("one_probe_client_time", "Post request Client timeout").Default("10s").Duration()
	ProbeURL = kingpin.Flag("prober", "This is server for probe").Default("http://21.64.64.200:8018/v1/hc/network/tcp/portstatus").String()
)


func ConfigParse() {
	kingpin.Version(version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
}