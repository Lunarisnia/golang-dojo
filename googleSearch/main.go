package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Search interface {
	Fetch() []string
}

type Web struct {
	query string
}
type Images struct {
	query string
}
type Videos struct {
	query string
}

func (w Web) Fetch() []string {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return fakeSearch.results["websites"]
}
func (i Images) Fetch() []string {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return fakeSearch.results["images"]
}
func (v Videos) Fetch() []string {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return fakeSearch.results["videos"]
}

// This code launch the services indepedently with goroutine
// And multiple instances is made using replicas to ensure robustnes and speed
// Because out of all the instances the first one to be ready will be served
func Google(query string) (results []string) {
	c := make(chan []string)
	web1 := Web{query: query}
	web2 := Web{query: query}
	web3 := Web{query: query}
	web4 := Web{query: query}
	web5 := Web{query: query}
	img1 := Images{query: query}
	img2 := Images{query: query}
	img3 := Images{query: query}
	img4 := Images{query: query}
	img5 := Images{query: query}
	vid1 := Videos{query: query}
	vid2 := Videos{query: query}
	vid3 := Videos{query: query}
	vid4 := Videos{query: query}
	vid5 := Videos{query: query}
	go func() { c <- Replicas(web1, web2, web3, web4, web5) }()
	go func() { c <- Replicas(img1, img2, img3, img4, img5) }()
	go func() { c <- Replicas(vid1, vid2, vid3, vid4, vid5) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case res := <-c:
			results = append(results, res...)
		case <-timeout:
			fmt.Println("Connection timeout!")
			return
		}
	}
	return
}

func Replicas(search ...Search) []string {
	c := make(chan []string)
	searchReplica := func(i int) { c <- search[i].Fetch() }
	for i := range search {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	start := time.Now()
	query := "Golang is awesome!"
	result := Google(query)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(result)
}

type fakeResult struct {
	results map[string][]string
}

var fakeSearch = fakeResult{
	results: map[string][]string{
		"images":   {"Image1", "Image2", "Image3"},
		"websites": {"Website1", "Website2", "Website3"},
		"videos":   {"Videos1", "Videos2", "Videos3"},
	},
}
