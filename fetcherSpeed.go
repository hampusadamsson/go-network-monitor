package main

import (
	"fmt"
	"log"

	"github.com/showwin/speedtest-go/speedtest"
)

type RealFetcher struct {
	down float64
	up   float64
}

func (df *RealFetcher) GetDown() float64 {
	return df.down
}

func (df *RealFetcher) GetUp() float64 {
	return df.up
}

func (df *RealFetcher) Eval() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
			df.down = 0.0
			df.up = 0.0
		}
	}()

	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})
	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(true)
		s.UploadTest(true)
		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		df.down = s.DLSpeed
		df.up = s.ULSpeed
	}

}
