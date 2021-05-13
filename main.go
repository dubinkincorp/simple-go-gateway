package main

import (
	"fmt"
	"gateway-go/handlers"
	"gateway-go/middleware"
	"net/http"
)

func main() {
	fmt.Println("Start server")

	http.Handle("/pivo", middleware.InterceptJWT(handlers.GreetingHandler))
	http.ListenAndServe(":8080", nil)
}