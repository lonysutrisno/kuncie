package main

import (
	"log"
	"net/http"

	"github.com/lonysutrisno/kuncie/pkg"
	"github.com/lonysutrisno/kuncie/service/checkout/delivery"
)

func main() {
	//init kuncie
	pkg.Up()
	router := delivery.NewRouter()
	log.Println("ListenAndServe kuncie-svc:8082")
	err := http.ListenAndServe(":8082", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
