package main

import (
	"log"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d) // решил реализовать через after которое будет ждать какое то время и потом засыпать
}

func main() {
	log.Println("Старт")
	sleep(5 * time.Second)
	log.Println("Конец")
}
