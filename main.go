package main

import (
	"fmt"
	"strconv"
	"time"

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
		userInput = inputInt()
		for userInput < 1 || userInput > 6 {
			fmt.Println("Вы ввели неверное значение. Выберите из списка (1-6):")
			userInput = inputInt()
		}
		switch userInput {
		case 1:
			var flowInput int
			fmt.Println("Выберите из списка:")
			fmt.Println("1. Divided. Эта цункция выводит числа кратные трём (от 1 до 100)")
			fmt.Println("2. FizzBuzz. Эта функция выводит числа кратные трём (Fizz) и пяти (Buzz)(от 1 до 100)")
			flowInput = inputInt()
			for flowInput < 1 || flowInput > 2 {
				fmt.Println("Вы ввели неверное значение. Выберите из списка (1-2):")
				flowInput = inputInt()

			}
			switch flowInput {
			case 1:
				fmt.Println("Результат:")
				quests.Divided()
			case 2:
				fmt.Println("Результат:")
				quests.FizzBuzz()
			}
		case 2:
			var funcInput int
			fmt.Println("Выберите из списка:")
			fmt.Println("1. Max. Эта функция выводит максимальное целое число из всех переданых в функцию")
			fmt.Println("2. EvenOdd. Эта функция определяет, чётное или нечётное число передано в функцию и возвращает кратность двум.")
			fmt.Println("3. MakeOddGenerator. Эта функция возвращает функцию, которая генерирует нечётное число")
			fmt.Println("4. Fib. Последовательность Фиббхуесоса.")
			funcInput = inputInt()
			for funcInput < 1 || funcInput > 4 {
				fmt.Println("Вы ввели неверное значение. Выберите из списка (1-4):")
				funcInput = inputInt()

			}
			switch funcInput {
			case 1:
				fmt.Println("func Max(nums ...int) int")
				fmt.Println("Введите числа для передачи их в функцию. Каждое число вводится отдельно (0 для выхода из цикла):")
				var nums []int
				for {
					var num int
					fmt.Scanln(&num)
					if num == 0 {
						break
					}
					nums = append(nums, num)
				}
				out, err := quests.Max(nums...)
				if err != nil {
					fmt.Println("Ошибка: ", err)
					break
				}
				fmt.Println("Результат: ", out)
			case 2:
				fmt.Println("func EvenOdd(a int) (int, bool)")
				fmt.Println("Введите число для передачи в функцию:")
				var a int
				var evenOdd bool
				fmt.Scanln(&a)
				a, evenOdd = quests.EvenOdd(a)
				fmt.Println("Результат: ", a, evenOdd)
			case 3:
				var g int = 1
				fmt.Println("func MakeOddGenerator() func() uint")
				fmt.Println("Данная функция генерирует нечётные числа. Каждый раз при вызове этой функции, генерируется следующее нечётное число")
				for {
					fmt.Println("Введите количество нечётных чисел")
					var num int
					fmt.Scanln(&num)
					for i := 0; i < num; i++ {
						oddGenerator := quests.MakeOddGenerator(&g)
						fmt.Println("Результат: ", oddGenerator())
					}
					var makeOddGeneratorInput string = ""
					fmt.Println("Для повтора нажмите enter (при этом будут генерироваться новые числа), чтобы выйти из цикла введите любой символ:")
					fmt.Scanln(&makeOddGeneratorInput)
					if makeOddGeneratorInput != "" {
						break
					}
				}
			case 4:
				fmt.Println("func Fib(n int) int")
				fmt.Println("Введите число для передачи в функцию:")
				var n int
				fmt.Scanln(&n)
				fmt.Println("Результат: ", quests.Fib(n))
			}
		case 3:
			fmt.Println("1. Min. Выводит минимальное число в массиве")
			fmt.Println("2. Max. Выводит максимальное число в заданном массиве")
			fmt.Println("Выберите из списка:")
			var massInput int
			massInput = inputInt()
			for massInput < 1 || massInput > 2 {
				fmt.Println("Вы ввели неверное значение. Выберите из списка (1-2):")
				massInput = inputInt()

			}
			switch massInput {
			case 1:
				fmt.Println("func Min(x []int) int")
				fmt.Println("Введите массив. Каждое число вводится отдельно. Для выхода из цикла, введите 0:")
				var mass []int
				for {
					var num int
					fmt.Scanln(&num)
					if num == 0 {
						break
					}
					mass = append(mass, num)
				}
				fmt.Println("Результат: ", quests.Min(mass))
			case 2:
				fmt.Println("func MinComp()")
				quests.MinComp()
			}
		case 4:
			fmt.Println("Определение площади и периметра фигур (Circle и Rectangle)")
			fmt.Println("Введите радиус окружности:")
			var r float64
			r = inputFloat64()
			circle := quests.Circle{Radius: r}
			fmt.Println("Введите первую сторону прямоугольника:")
			var a float64
			a = inputFloat64()
			fmt.Println("Введите вторую сторону прямоугольника:")
			var b float64
			b = inputFloat64()
			rectangle := quests.Rectangle{Width: a, Height: b}
			fmt.Println("Результат:")
			fmt.Println("Площадь окружности: ", circle.Area())
			fmt.Println("Периметр окружности: ", circle.Perimeter())
			fmt.Println("Площадь прямоугольника: ", rectangle.Area())
			fmt.Println("Периметр прямоугольника: ", rectangle.Perimeter())
		case 5:
			fmt.Println("func Sleep(duration time.Duration)")
			fmt.Println("Функция ожидания (через time.After). Введите число секунд до завершения функции:")
			var sec int
			sec = inputInt()
			duration := time.Duration(sec) * time.Second
			fmt.Println("Начало ожидания...")
			quests.Sleep(duration)
			fmt.Println("Конец ожидания.")
		case 6:
			fmt.Println("func Swap(x *int, y *int) (*int, *int)")
			fmt.Println("Данная функция меняет местами две переменные")
			fmt.Println("Введите первое значение в переменную a:")
			var a int
			a = inputInt()
			fmt.Println("Введите второе значение в переменную b:")
			var b int
			b = inputInt()
			fmt.Println("a = ", a, "b = ", b)
			quests.Swap(&a, &b)
			fmt.Println("Результат: a = ", a, ", b = ", b)
		}
	}
}

// Проверка введённых пользователем данных. int
func inputInt() int {
	var userInput string
	var num int
	var err error
	fmt.Println("Введите целое число:")
	for {
		fmt.Scanln(&userInput)

		num, err = strconv.Atoi(userInput)
		if err == nil {
			break
		} else {
			fmt.Println("Ошибка: введите целое число.")
		}
	}
	return num
}

// Проверка введённых пользователем данных. float64
func inputFloat64() float64 {
	var userInput string
	var num float64
	var err error
	fmt.Println("Введите целое число:")
	for {
		fmt.Scanln(&userInput)

		num, err = strconv.ParseFloat(userInput, 5)
		if err == nil {
			break
		} else {
			fmt.Println("Ошибка: введите целое число.")
		}
	}
	return num
}
