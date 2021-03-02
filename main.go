package main

import (
	"math"
	"strconv"
	"time"
)

func backgroundTask() {
	c := newConfig()
	ticker := time.NewTicker(time.Duration(c.RequestInterval) * time.Second)
	for _ = range ticker.C {
		for _, share := range c.Shares {
			a := getCurrentPrice(share.Ticker, *c)
			profit := (a / share.Startprice * 100) - 100
			if profit > 5 {
				profitRate := math.Round(profit*100) / 100
				//fmt.Println(share.Ticker, math.Round(profit*100)/100)
				sendMessage(share.Ticker+": +"+strconv.FormatFloat(profitRate, 'f', 2, 64)+"%", *c)
			}
		}
	}
}

func main() {
	go backgroundTask()
	select {}

}
