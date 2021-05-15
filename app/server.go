package main

import (
	"net/http"
	"cryptocur/controllers"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    data := []byte("Hello World!") // slice of bytes

    res.Write(data)
}

func main() {
    // handler := HttpHandler{}
    // listen and serve
	// http.ListenAndServe(":9000", handler)
	controllers.FetchData()
}