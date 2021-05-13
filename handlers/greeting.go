package handlers

import "net/http"

var GreetingHandler = func(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello"))
}