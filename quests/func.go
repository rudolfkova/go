package quests

import "fmt"

//Нахождение максимального значения в массиве
func Max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("массив чисел пуст")
	}
	g := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if g < nums[i] {
			g = nums[i]
		}
	}
	return g, nil
}

//Определение чётности/нечётности, с выводом int и bool
func EvenOdd(a int) (int, bool) {
	if a%2 == 0 {
		return a / 2, true
	} else {
		return a / 2, false
	}
}

//Генератор нечётных чисел
func MakeOddGenerator(g *int) func() int {
	return func() (ret int) {
		ret = *g
		*g += 1
		if *g%2 == 0 {
			*g += 1
		}
		return
	}
}

//Последовательность Фиббаначи
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
