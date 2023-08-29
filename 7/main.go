package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	mutex := sync.Mutex{}

	m := make(map[int]int) // Создание map

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) { // В анонимной горутине происходит доб-е чисел в map
			defer wg.Done()
			mutex.Lock() // Использую мьютекс чтобы позволить каждой горутине произвести записьв map
			m[i] = i * i
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)

}
