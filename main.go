package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", homeHandler)
	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); nil != err {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
