package probe

import (
	"I2Oprobe/internal/g"
	"I2Oprobe/internal/log"
	"I2Oprobe/internal/model"
	"fmt"
	"sync"
	"time"
)

type (
	App interface {
		Post(data interface{})
	}
)


type Server struct {
	data []interface{}
	appCli *appProbe
}

func NewServer(data []interface{}, appcli *appProbe) *Server {
	return &Server{
		data: data,
		appCli: appcli,
	}
}

func (s *Server) Run() {
	s.Handler()
}

//
//@description concurrency of post
//
func (s *Server) Handler() {
	goConCur := NewGoConCur()
	wg := &sync.WaitGroup{}

	for i :=0; i < len(s.data); i++{  // don`t use `FOR...RANGE`. Because of valueIndex.

		wg.Add(1)
		task := s.data[i]
		// IP relate with SHOPNAME
		g.MetricsIpShop[task.(model.Source).Address] = task.(model.Source).Shopname
		gofunc := func() {
			//time.Sleep(1 * time.Second)
			s.appCli.Post(task)
			wg.Done()
		}
		goConCur.Run(gofunc)
	}

	wg.Wait()
	log.Info(fmt.Sprintf("END at %v", time.Now()))
	fmt.Println("END")


	//res := func(k, v interface{}) bool{
	//	g.MetricsMapCache.Store(k.(string), v.([]model.MetricsOpts))
	//	return true
	//}
	//g.MetricsMap.Range(res)
	//
	//resBad := func(k,v interface{})bool {
	//	g.MetricsReqMapCache.Store(k.(string), v.(string))
	//	return true
	//}
	//g.MetricsReqMapCache.Range(resBad)

}