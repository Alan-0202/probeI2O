package list

import (
	"I2Oprobe/internal/g"
)

type SourceFile interface {
	Read()
	ToChan() []interface{}
}

type Dolist struct {}

func NewDoList() *Dolist {
	return &Dolist{}
}

func (dl *Dolist) Run() []interface{} {
	collect := NewSourFile(*g.SourcePath)
	collect.Read()
	return collect.ToChan()
}
