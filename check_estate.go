package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func callAPI() error {
	url := "https://rest.immobilienscout24.de/restapi/api/search/v1.0/search/radius?realestatetype=ApartmentRent&geocoordinates=52.512303;13.431191;1"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	username := os.Getenv("IM24_USERNAME")
	password := os.Getenv("IM24_PASSWORD")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	s := string(bodyText)
	fmt.Println(s)
	return nil
}
