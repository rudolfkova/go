package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало программы")
	Sleep(5 * time.Second)
	fmt.Println("Конец программы")
}
