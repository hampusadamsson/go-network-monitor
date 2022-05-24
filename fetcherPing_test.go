package main

import (
	"log"
	"testing"
)

func TestPinger(t *testing.T) {
	p := PingFetcher{addr: "www.google.com"}
	p.Eval()
	log.Println(p)
}
