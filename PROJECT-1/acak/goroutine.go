package main

import (
	"fmt"
	"time"
)

func main() {
	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	for i := 0; i < 8; i++ {
		id := i % 2
		go func() {
			if id == 1 {
				fmt.Println(data1, id+1)
			} else {
				fmt.Println(data2, id+1)
			}
		}()
	}

	time.Sleep(1 * time.Second)
}
