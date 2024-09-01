package main

import (
	"fmt"
	"github.com/rudolfkova/go/quests"
	"time"
)

func main() {
	fmt.Println("Начало программы")
	quests.Sleep(5 * time.Second)
	fmt.Println("Конец программы")
}
