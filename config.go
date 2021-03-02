package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	MoexAuthentication string `yaml:"moexAuthentication"`
	MoexPrefixHttps    string `yaml:"moexPrefixHttps"`
	Password           string `yaml:"password"`
	UserName           string `yaml:"userName"`
	RequestInterval    int16  `yaml:"requestInterval"`
	TlgChatId          int64  `yaml:"tlgChatId"`
	TlgAccessToken     string `yaml:"tlgAccessToken"`
	Shares             []struct {
		Ticker     string  `yaml:"ticker"`
		Startprice float64 `yaml:"startprice"`
	} `yaml:"shares"`
}

func newConfig() *config {

	c := &config{}

	file, err := ioutil.ReadFile(".vars.yml")
	if err != nil {
		log.Printf("file.Get err #%v", err)
	}
	err = yaml.Unmarshal(file, c)

	if err != nil {
		log.Fatalf("Unmarshall: %v", err)
	}

	return c

}
