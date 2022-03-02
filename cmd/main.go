package main

import (
	"I2Oprobe/internal/g"
	"I2Oprobe/internal/probe"
	"I2Oprobe/web"
	"time"
)

func main() {
	g.ConfigParse()


	// crontal to collect data
	go func() {
		probe.Do()
		ticker := time.NewTicker(*g.ProbeRangeTime)
		for {
			select {
			case <- ticker.C:
				probe.Do()
			}
		}
	}()

	web.Start()
}
