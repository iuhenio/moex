package main

import (
	"fmt"
	"math"
	"time"
)

func backgroundTask() {
	c := newConfig()
	ticker := time.NewTicker(time.Duration(c.RequestInterval) * time.Second)
	for _ = range ticker.C {
		for _, share := range c.Shares {
			a := getCurrentPrice(share.Ticker)
			profit := (a / share.Startprice * 100) - 100
			if profit > 5 {
				fmt.Println(share.Ticker, math.Round(profit*100)/100)
			}
		}
	}
}

func main() {
	go backgroundTask()
	select {}

}
