package main

import (
	"fmt"
	"net/http"

	router "WebApp/src"

	"github.com/gorilla/mux"
)

var port = ":1337"

func main() {
	r := mux.NewRouter()
	router.SetRouter(r) //Setting website Path

	fmt.Println("Server Listen at: " + port)
	http.ListenAndServe(port, r)
	return
}
