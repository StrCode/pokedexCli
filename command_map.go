package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const POKEAPIURL = "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"

type Config struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(conf *Config) error {
	postUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	if len(conf.Next) != 0 {
		postUrl = conf.Next
	}

	req, err := http.NewRequest("GET", postUrl, nil)
	if err != nil {
		return fmt.Errorf("could not create request:%v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not get response:%v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
	}

	if err := json.Unmarshal(data, conf); err != nil {
		return fmt.Errorf("could not marshall data into config:%v", err)
	}

	for _, area := range conf.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(conf *Config) error {
	postUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	if conf.Previous != nil {
		postUrl = *conf.Previous
		fmt.Println("get people here" + postUrl)
	}

	req, err := http.NewRequest("GET", postUrl, nil)
	if err != nil {
		return fmt.Errorf("could not create request:%v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not get response:%v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
	}

	if err := json.Unmarshal(data, conf); err != nil {
		return fmt.Errorf("could not marshall data into config:%v", err)
	}

	for _, area := range conf.Results {
		fmt.Println(area.Name)
	}

	return nil
}
