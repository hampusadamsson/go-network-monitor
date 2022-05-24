package main

import (
	"fmt"
)

type DummyFetcher struct {
	i    int
	down float64
	up   float64
}

func (df *DummyFetcher) GetDown() float64 {
	return df.down
}

func (df *DummyFetcher) GetUp() float64 {
	return df.up
}

func (df *DummyFetcher) Eval() {
	defer func() {
		if r := recover(); r != nil {
			df.down = 0.0
			df.up = 0.0
		}
	}()
	df.i++
	fmt.Println(df.i)
	if df.i == 3 {
		panic("ERROR")
	}
	df.down = 10.0
	df.up = 2.0
}
