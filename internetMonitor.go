package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"
)

type InternetSpeedMonitor struct {
	logFilename      string
	timeBetweenPolls time.Duration
	fetcher          fetcher
	datapointsToRead int
}

func CreateSpeedMonitor() InternetSpeedMonitor {
	ttl := time.Minute * 30
	tmp := RealFetcher{}
	ism := InternetSpeedMonitor{logFilename: "log/data.log", timeBetweenPolls: ttl, fetcher: &tmp, datapointsToRead: 15}
	return ism
}

func CreatePingMonitor() InternetSpeedMonitor {
	ttl := time.Second * 30
	tmp := PingFetcher{addr: "www.google.com"}
	ism := InternetSpeedMonitor{logFilename: "log/dataping.log", timeBetweenPolls: ttl, fetcher: &tmp, datapointsToRead: 15}
	return ism
}

func CreateDummy() InternetSpeedMonitor {
	ttl := time.Second * 1
	tmp := DummyFetcher{}
	ism := InternetSpeedMonitor{logFilename: "log/data-dummy.log", timeBetweenPolls: ttl, fetcher: &tmp, datapointsToRead: 5}
	return ism
}

func (ism InternetSpeedMonitor) start() {
	for {
		dataPoint := ism.fetchNetworkSpeed()
		ism.save(dataPoint)
		time.Sleep(ism.timeBetweenPolls)
	}
}

func (ism InternetSpeedMonitor) fetchNetworkSpeed() DataPoint {
	now := time.Now().String()[:19]
	ism.fetcher.Eval()
	return DataPoint{Time: now, Download: ism.fetcher.GetDown(), Upload: ism.fetcher.GetUp()}
}

func (ism InternetSpeedMonitor) save(dataPoint DataPoint) {
	data, _ := json.Marshal(dataPoint)
	log.Println(string(data))
	f, err := os.OpenFile(ism.logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(string(data) + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (ism InternetSpeedMonitor) read() []DataPoint {
	file, err := os.Open(ism.logFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var dataPoint DataPoint
	var stack []DataPoint
	for scanner.Scan() {
		json.Unmarshal([]byte(scanner.Text()), &dataPoint)
		stack = append(stack, dataPoint) //push
		if len(stack) > ism.datapointsToRead {
			stack = stack[1:] // Pop
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stack
}
