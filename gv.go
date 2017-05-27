package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func geturl() string {

	apiurl := "https://godville.net/gods/api/"
	//_, keyID, vCode, _, _, _ := config() "https://godville.net/gods/api/isnog00d/e8babf8ed27d"
	_, god, key, _, _ := config()

	jsonurl := apiurl + god + "/" + key

	//fmt.Println(xmlurl)
	return jsonurl
}

func vaterpaz() {

	//Url := "https://godville.net/gods/api/isnog00d/e8babf8ed27d"
	Url := geturl()
	res, err := http.Get(Url)
	if err != nil {
		fmt.Println(err)
		return

	}

	gvdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	res.Body.Close()

	type Result struct {
		Bag   int `json:"inventory_max_num"`
		Items int `json:"inventory_num"`
	}

	var result Result
	json.Unmarshal(gvdata, &result)
	fmt.Printf("Bag = %v", result.Bag)
	fmt.Println("")
	fmt.Printf("items = %v", result.Items)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(result.Bag - result.Items)

	if result.Bag-result.Items < 5 {
		//fmt.Println("Rich hero!")
		achtung()

	}
}
