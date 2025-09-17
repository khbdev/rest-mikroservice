package main

import (
	"apiGetWay/router"
	"log"
)


func main() {
	r := router.SetupRouter()
	log.Println("API Gateway running on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}