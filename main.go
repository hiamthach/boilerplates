package main

import (
	"fmt"
	app "go-gin/app"
	apphost "go-gin/app/apphost"
)

func init() {
	fmt.Println("Initialize app...")
	apphost.IoC()
}

func main() {
	// start gin server
	app.Start(&apphost.AppConfig, apphost.AppLog)
}
