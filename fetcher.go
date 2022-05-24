package main

type fetcher interface {
	Eval()
	GetDown() float64
	GetUp() float64
}
