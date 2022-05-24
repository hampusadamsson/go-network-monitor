package main

import (
	"log"
	"time"

	"github.com/go-ping/ping"
)

type PingFetcher struct {
	avgPing      float64
	stdDeviation float64
	addr         string
}

func (df *PingFetcher) GetDown() float64 {
	return df.avgPing
}

func (df *PingFetcher) GetUp() float64 {
	return df.stdDeviation
}

func (df *PingFetcher) Eval() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
			df.avgPing = 0.0
			df.stdDeviation = 0.0
		}
	}()

	pinger, err := ping.NewPinger(df.addr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	pinger.Timeout = time.Second * 30
	pinger.Count = 5
	err = pinger.Run()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	df.avgPing = float64(stats.AvgRtt.Milliseconds())
	df.stdDeviation = float64(stats.StdDevRtt.Milliseconds())
	log.Println(stats)
}
