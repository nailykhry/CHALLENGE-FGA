package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	h8HelperRand "github.com/novalagung/gubrak"
)

type Data struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	for {
		var randomNumberWater int = h8HelperRand.RandomInt(0, 100)
		var randomNumberWind int = h8HelperRand.RandomInt(0, 100)
		sendPost(randomNumberWater, randomNumberWind)
		time.Sleep(15 * time.Second)
	}
}

func getStatusWater(water int) string {
	if water < 5 {
		return "aman"
	} else if water >= 5 && water <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getStatusWind(wind int) string {
	if wind < 6 {
		return "aman"
	} else if wind >= 6 && wind <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func sendPost(randomNumberWater int, randomNumberWind int) {
	// Tentukan status untuk water dan wind
	waterStatus := getStatusWater(randomNumberWater)
	windStatus := getStatusWind(randomNumberWind)

	// Buat data dalam format JSON
	data := map[string]interface{}{
		"water": randomNumberWater,
		"wind":  randomNumberWind,
	}
	requestJSON, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		fmt.Println(err)
		return
	}

	url := "https://jsonplaceholder.typicode.com/posts"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(string(body))
	fmt.Printf("status water : %v \nstatus wind : %v\n\n\n", waterStatus, windStatus)
}
