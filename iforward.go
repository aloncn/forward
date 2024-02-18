package main

import (
	"fmt"
	"iforward/internal/service"
	"log"
	"net/http"
)

const addr = "0.0.0.0:9009"

func main() {
	http.HandleFunc("/", service.GetDo)
	http.HandleFunc("/proxy", service.ProxyDo)

	log.Println("Start server at ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("http listen failed")
	}

}
