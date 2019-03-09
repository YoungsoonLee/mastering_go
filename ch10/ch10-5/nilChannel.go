package main

import (
	"fmt"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.c:
			c = nil
			fmt.Println(sum)
		}
	}
}
