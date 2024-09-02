package main

import (
	"fmt"

	"github.com/rudolfkova/tutor/quests"
)

func main() {
	var out string = ""
	for out == "" {
		fmt.Println("Чтобы начать программу, нажмите Enter. Для выхода, введите любой символ:")
		fmt.Scanln(&out)
		if out != "" {
			break
		}
		var userInput int
		fmt.Println("Выберите из списка:")
		fmt.Println("1. Управление потоком (flow)")
		fmt.Println("2. Функции (func)")
		fmt.Println("3. Массивы (mass)")
		fmt.Println("4. Структуры и интерфейсы (shape)")
		fmt.Println("5. Потоки (sleep)")
		fmt.Println("6. Указатели (swap)")
		fmt.Scanln(&userInput)
		if userInput < 1 || userInput > 6 {
			fmt.Println("Вы ввели неверное значение. Выберите из списка (1-6):")
			fmt.Scanln(&userInput)
		}
		switch userInput {
		case 1:
			var flowInput int
			fmt.Println("Выберите из списка:")
			fmt.Println("1. Divided. Эта цункция выводит числа кратные трём (от 1 до 100)")
			fmt.Println("2. FizzBuzz. Эта функция выводит числа кратные трём (Fizz) и пяти (Buzz)(от 1 до 100)")
			fmt.Scanln(&flowInput)
			switch flowInput {
			case 1:
				quests.Divided()
			case 2:
				quests.FizzBuzz()
			}
		case 2:
			var funcInput int
			fmt.Println("Выберите из списка:")
			fmt.Println("1. Max. Эта функция выводит максимальное целое число из всех переданых в функцию")
			fmt.Println("2. EvenOdd. Эта функция определяет, чётное или нечётное число передано в функцию и возвращает кратность двум.")
			fmt.Println("3. MakeOddGenerator. Эта функция возвращает функцию, которая генерирует нечётное число")
			fmt.Println("4. Fib. Последовательность Фиббхуесоса.")
			fmt.Scanln(&funcInput)
			switch funcInput {
			case 1:
				var nums []int
				fmt.Println("Введите числа для передачи их в функцию. Каждое число вводится отдельно (0 для выхода из цикла):")
				for {
					var num int
					fmt.Scanln(&num)
					if num == 0 {
						break
					}
					nums = append(nums, num)
				}
				fmt.Println(quests.Max(nums...))
			}

		}
	}
}
