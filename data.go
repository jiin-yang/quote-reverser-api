package main

type ResponseServer struct {
	Author	string	`json:"author"`
	Text	[]string	`json:"text"`
}

var ResponseServerData []ResponseServer