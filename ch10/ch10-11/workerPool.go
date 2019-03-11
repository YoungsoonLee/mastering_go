package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int // request
}

type Data struct {
	job    Client
	square int // result of request
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients { // read clients channel
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	go create(nJobs) // create clients
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\t int: ", d.job.id)
			fmt.Printf("%d\t squre: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()
	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}
