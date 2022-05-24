package main

import (
	"github.com/kaimu/speedtest"
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
			df.down = 0.0
			df.up = 0.0
		}
	}()
	down, up, _ := speedtest.Ookla()
	df.down = down
	df.up = up
}
