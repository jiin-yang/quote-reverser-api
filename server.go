package main

import (
	"encoding/json"
	"net/http"
)

type server struct {

}

func (s *server) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := GetDataClient()
	w.Write(response)
}

func GetDataClient() []byte {

	responseData := SendRequest()

	DataServerFormat := make(map[string][]string)

	for i := 0; i < len(responseData); i++ {
		DataServerFormat[responseData[i].Author] = append(DataServerFormat[responseData[i].Author], reverseString(responseData[i].Text))
	}

	responseServerArray := make([]ResponseServer, 0)
	
	for key, value := range DataServerFormat{
		tempResponseServerObj := ResponseServer{
			Author: key,
			Quote:  value,
		}

		responseServerArray = append(responseServerArray, tempResponseServerObj)
	}
	DataJSONFormat, _ := json.MarshalIndent(responseServerArray, "", "  ")

	return DataJSONFormat
}

func reverseString(str string) string{
	byte_str := []rune(str)
	var res []rune
	for i := len(byte_str) - 1; i >= 0; i-- {
		res = append(res , byte_str[i])
	}
	return string(res)
}