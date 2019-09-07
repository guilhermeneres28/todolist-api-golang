package main

import (
	"api/api"
	"fmt"
)

func main() {
	channel := make(chan string)
	go api.RunWebServer(channel)
	fmt.Println(<-channel)
}
