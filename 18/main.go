package main

import (
	"fmt"
	"sync"
)

type Counter struct { // структура с одним полем - числом, которое изначально равно 0 и будет инкрементироваться
	num int
}

func (c *Counter) Increment() { // простой метод инкреминации
	c.num++
}

func main() {
	mutex := sync.Mutex{}
	var wg sync.WaitGroup
	count := Counter{}

	for i := 0; i < 10; i++ { // В цикле вызываю метод инкреминации
		wg.Add(1)
		go func() { // каждая анонимная горутина берет на себя число и инкреминирует его
			mutex.Lock()
			count.Increment() // вызываю внутри горутины метод инкреминации
			mutex.Unlock()
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(count.num)
}
