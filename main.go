package main

import "sync"

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	d := CreateSpeedMonitor()
	go d.start()

	pm := CreatePingMonitor()
	go pm.start()

	RunServer()

	wg.Wait()

}
