package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Stock struct {
	Marketdata Marketdata `json:"marketdata"`
}

type Marketdata struct {
	Columns []string    `json:"columns"`
	Data    [][]float64 `json:"data"`
}

func basicAuth(c *config) string {
	client := &http.Client{}
	url := c.MoexAuthentication
	req, err := http.NewRequest("GET", url, nil)
	auth := c.UserName + ":" + c.Password
	credentials := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+credentials)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	cookie := resp.Cookies()
	fmt.Println(cookie)
	respString := string(bodyText)
	return respString

}

func getCurrentPrice(s string) float64 {
	client := &http.Client{}
	url := "https://iss.moex.com/iss/engines/stock/markets/shares/boards/tqbr/securities/" + s + ".json?iss.meta=off&iss.data=on&iss.json=compact&iss.only=marketdata&marketdata.columns=LAST"
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyText))
	var stock Stock
	json.Unmarshal([]byte(bodyText), &stock)
	//cookie := resp.Cookies()
	//fmt.Println(cookie)

	return stock.Marketdata.Data[0][0]

}
