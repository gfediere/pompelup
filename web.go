package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type settings struct {
	Id       int    `json:"id"`
	Product  string `json:"product"`
	Quantity int    `json:"Quantity"`
	Time     struct {
		Hour   int `json:"Hour"`
		Minute int `json:"Minute"`
	} `json:"Time"`
}

func getSettings() []settings {
	raw, err := ioutil.ReadFile("./parameters.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var c []settings
	json.Unmarshal(raw, &c)
	return c
}

func main() {

	pages := getSettings()
	for _, p := range pages {
		fmt.Printf("ID: %d, Product: %s, Quantity: %d, Time { Hour: %d, Minutes: %d }\n", p.Id, p.Product, p.Quantity, p.Time.Hour, p.Time.Minute)
	}
}
