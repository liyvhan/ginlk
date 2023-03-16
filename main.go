package main

import (
	"my_webserver/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: core,
	}
	server.ListenAndServe()
}
