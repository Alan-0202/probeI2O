package probe

import "I2Oprobe/internal/g"

type GoConCur struct {
	num int
	c chan struct{}
}

func NewGoConCur() *GoConCur {
	return &GoConCur{
		num: *g.ConcurNum,
		c: make(chan struct{}, *g.ConcurNum),
	}
}

func (g *GoConCur) Run( f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<- g.c
	}()
}