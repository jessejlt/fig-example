package main

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

func main() {

	http.HandleFunc("/", homeHandler)
	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); nil != err {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	c, err := redis.Dial("tcp", "redis_1:6379")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer c.Close()

	n, err := redis.Int(c.Do("INCR", "cons"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Connections: %d\n", n)
}
