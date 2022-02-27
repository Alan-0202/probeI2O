package probe

import (
	"I2Oprobe/internal/g"
	"I2Oprobe/internal/model"
	"fmt"
	"sync"
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
			fmt.Printf("Consumer: %d\n", task.(model.Source).Shopname)
			//time.Sleep(1 * time.Second)
			s.appCli.Post(task)
			wg.Done()
		}
		goConCur.Run(gofunc)
	}

	wg.Wait()
	fmt.Println("END")
}