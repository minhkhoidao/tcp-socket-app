package main

import (
	"fmt"
	"sync"
)

// TaskA simulates a task that produces data
func TaskA(dataChan chan<- string) {
	// Simulate data production
	data := "Data from Task A"
	// Send data to Task B
	dataChan <- data
	close(dataChan) // Close the channel once data is sent
}

// TaskB simulates a task that consumes and processes data
func TaskB(dataChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Receive data from Task A
	data, ok := <-dataChan
	if !ok {
		fmt.Println("Channel closed before receiving data")
		return
	}
	// Process the data
	fmt.Printf("Task B received: %s\n", data)
	// Simulate data processing
}

func main() {
	var wg sync.WaitGroup
	dataChan := make(chan string)

	wg.Add(1) // Increment the WaitGroup counter

	go TaskA(dataChan)      // Start Task A in a goroutine
	go TaskB(dataChan, &wg) // Start Task B in a goroutine

	wg.Wait() // Wait for all goroutines to finish
}
