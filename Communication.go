package main

import "fmt"

func worker(ch chan string) {
	ch <- "Task completed"
}

func main() {
	ch := make(chan string)

	go worker(ch)

	msg := <-ch
	fmt.Println(msg)
}
