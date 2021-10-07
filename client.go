package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseClient struct {
	Author	string	`json:"author"`
	Text	string	`json:"text"`
}

func SendRequest() []ResponseClient {

	response, err := http.Get("https://type.fit/api/quotes")

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []ResponseClient
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}