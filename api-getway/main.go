package main

import (
	"apiGetWay/router"
	"log"
)


func main() {
	r := router.SetupRouter()
	log.Println("API Gateway running on :8086")
	if err := r.Run(":8086"); err != nil {
		log.Fatal(err)
	}
}