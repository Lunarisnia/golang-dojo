package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Greeting string

func fakeFetchData() []string {
	words := []string{"LOREM", "IPSUM", "DOLOR", "SIT AMET", "BRUH"}
	response := make([]string, 10)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	for i := 0; i < 10; i++ {
		response[i] = words[rand.Intn(len(words))]
	}
	return response
}

func (g Greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// Wow holy shit goroutine is fast
	c := make(chan []string)
	for i := 0; i < 20; i++ {
		go func() { c <- fakeFetchData() }()
	}
	fmt.Fprint(w, <-c)
	
	// Wow holy shit goroutine is fast
	// resp := make([]string, 1)
	// for i := 0; i < 20; i++ {
	// 	resp = fakeFetchData()
	// }
	// resp := fakeFetchData()
	// fmt.Fprint(w, resp)
	fmt.Println(time.Since(start))
}

func main() {
	fmt.Println("Server is listening on: 4000")
	err := http.ListenAndServe("localhost:4000", Greeting("Hello, World!"))
	if err != nil {
		log.Fatal(err)
	}
}
