package main

import "github.com/ReneVallecillo/office/router"

func main() {
	router := router.InitRouter()
	router.Run(":8080")
}
