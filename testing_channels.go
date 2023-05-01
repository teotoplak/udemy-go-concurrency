package main

import "fmt"

func TestingChannels() {
	capacity := 2
	ch := make(chan string, capacity)
	go func() {
		ch <- "foo"
		ch <- "foo"
		ch <- "foo"
		ch <- "foo"
		close(ch)
	}()
	for {
		a, b := <-ch
		fmt.Printf("%s and %v \n", a, b)
		if b == false {
			return
		}
	}
}
