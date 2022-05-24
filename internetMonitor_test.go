package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// func TestInternetMonitor(t *testing.T) {
// 	ism := CreateDummy()
// 	ism.start()
// 	fmt.Println(ism)
// }

// func TestDataPoint(t *testing.T) {
// 	ism := CreateDummy()
// 	dp := ism.fetch()
// 	ism.save(dp)
// 	dp2 := ism.read()
// 	for v := range dp2 {
// 		log.Println(v)
// 	}
// }

func TestSerializeChan(t *testing.T) {
	ism := CreateDummy()
	dp := ism.read()
	fmt.Println(dp)
	jsonDatapoint, _ := json.Marshal(dp)
	fmt.Println(string(jsonDatapoint))
}
