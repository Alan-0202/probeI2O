package list

import (
	"I2Oprobe/internal/g"
	"testing"
)

func TestSourFile_ToChan(t *testing.T) {
	g.ConfigParse()
	dolist := new(Dolist)
	dolist.Run()
	t.Log(g.MetricsChan)
}
