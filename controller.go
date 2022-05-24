package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func getPing(w http.ResponseWriter, req *http.Request) {
	ism := CreatePingMonitor()
	dp := ism.read()
	jsonDatapoint, _ := json.Marshal(dp)
	fmt.Fprintf(w, string(jsonDatapoint))
}

func getSpeed(w http.ResponseWriter, req *http.Request) {
	ism := CreateSpeedMonitor()
	dp := ism.read()
	jsonDatapoint, _ := json.Marshal(dp)
	fmt.Fprintf(w, string(jsonDatapoint))
}

func ChanToSlice(ch interface{}) interface{} {
	chv := reflect.ValueOf(ch)
	slv := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(ch).Elem()), 0, 0)
	for {
		v, ok := chv.Recv()
		if !ok {
			return slv.Interface()
		}
		slv = reflect.Append(slv, v)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func RunServer() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/ping", getPing)
	http.HandleFunc("/speed", getSpeed)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
