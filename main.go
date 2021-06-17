package main

import (
	"log"
	"math"
	"net/http"
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
			if profit > float64(2*c.Commission) {
				profitRate := math.Round(profit*100) / 100
				moneyProfit := calcProfit(share.Number, share.Startprice, a, *c)
				if moneyProfit > c.MinProfit {
					sendMessage(share.Ticker+": +"+strconv.FormatFloat(profitRate, 'f', 2, 64)+"%, "+strconv.FormatFloat(moneyProfit, 'f', 2, 64), *c)
				}
				//fmt.Println(share.Ticker, math.Round(profit*100)/100)
			}
		}
	}
}

func calcProfit(number int16, startprice, sellprice float64, c config) float64 {
	profit := float64(number) * ((sellprice-startprice)*float64(100-c.Tax)/100 - (sellprice+startprice)*float64(c.Commission)/100)
	return profit
}

func startServer() {
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func main() {
	go startServer()
	go backgroundTask()
	select {}

}
