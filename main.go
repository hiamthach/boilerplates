package main

import (
	"go-microservices/app/router"
)

func main() {
	go router.StartGRPC()
	router.Start()
}
