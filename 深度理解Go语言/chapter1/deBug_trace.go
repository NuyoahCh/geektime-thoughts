package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello World")
	}
}
