package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// создаем горутину в которой будут воркеры выполнять какую либо работу
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	var c int
	fmt.Scan(&c)
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// описываем количество воркеров
	for w := 1; w <= c; w++ {
		go worker(w, jobs, results)
	}
	// описываем количество работы для воркеров
	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)
	// сигнал о том что работа сделана
	for a := 1; a <= 100; a++ {
		<-results
	}
}

// фукнция которая позволяет остановить выполнение работы программы по нажатию cntrl+C
func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}
