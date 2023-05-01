package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage(wg *sync.WaitGroup) {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func MutexExample() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	printMessage(&wg)
	wg.Wait()
	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	printMessage(&wg)
	wg.Wait()
	wg.Add(1)
	go updateMessage("Hello, world!")
	printMessage(&wg)
	wg.Wait()
}
