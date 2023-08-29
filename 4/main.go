package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	dataChan := make(chan string) // канал для хранения данных

	go func() { // анонимная горутина для записи данных в канал
		for {
			var data string
			fmt.Print("Введите данные: ")
			fmt.Scanln(&data)
			dataChan <- data
		}
	}()

	// Ожидаем нажатия Ctrl+C для завершения программы
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
	fmt.Println("Программа завершает работу")

	// Отправляем сигнал о завершении работы воркеров
	close(done)

	var wg sync.WaitGroup
	N := 5                   // кол-во воркеров
	for i := 0; i < N; i++ { // Запускаем N воркеров
		wg.Add(1)
		// Функция для чтения данных из канала и вывода в stdout
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case data := <-dataChan:
					fmt.Printf("Воркер %d получил данные: %s\n", workerID, data)
				case <-done:
					// Если получен сигнал о завершении работы, то выходим из цикла
					fmt.Printf("Воркер %d завершает работу\n", workerID)
					return
				}
			}
		}(i)
	}

	// Ожидаем завершения работы всех воркеров
	wg.Wait()
}
