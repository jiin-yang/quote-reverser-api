package main

type ResponseServer struct {
	Author	string		`json:"author"`
	Quote	[]string	`json:"quotes"`
}

var ResponseServerData []ResponseServer