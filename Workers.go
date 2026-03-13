package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Println("Worker", id, "started")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker", id, "finished")
}

func main() {
	for i := 1; i <= 5; i++ {
		go worker(i)
	}

	time.Sleep(2 * time.Second)
}
