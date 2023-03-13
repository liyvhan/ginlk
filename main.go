package main

import (
	"my_webserver/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: framework.NewCore(),
	}
	server.ListenAndServe()
}
