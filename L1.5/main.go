package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <N>")
		return
	}
	N, err := strconv.Atoi(os.Args[1])
	if err != nil || N < 0 {
		fmt.Println("Invalid N")
		return
	}
	ch := make(chan int)
	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()
	timeout := time.After(time.Duration(N) * time.Second)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-timeout:
			fmt.Println("Time is over")
			return
		}
	}
}
