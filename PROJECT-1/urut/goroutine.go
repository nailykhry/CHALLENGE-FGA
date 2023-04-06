package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	var w sync.WaitGroup
	var m sync.Mutex

	for i := 1; i <= 8; i++ {
		w.Add(1)
		go func(id int) {
			m.Lock()
			defer m.Unlock()
			if id%2 == 1 {
				fmt.Println(data1, id+1)
			} else {
				fmt.Println(data2, id+1)
			}
			w.Done()
		}(i % 2)
	}
	time.Sleep(1 * time.Second)
}
