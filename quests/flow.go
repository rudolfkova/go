package quests

import "fmt"

func Divided() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%5 == 0 || i%3 == 0 {
			if i%3 == 0 {
				fmt.Print("Fizz")
			}
			if i%5 == 0 {
				fmt.Print("Buzz")
			}
			fmt.Print(i)
			fmt.Println()
		}
	}
}
