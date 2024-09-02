package quests

//Нахождение максимального значения в массиве
func Max(nums ...int) int {
	g := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if g < nums[i] {
			g = nums[i]
		}
	}
	return g
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
func MakeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 1
		if i%2 == 0 {
			i += 1
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
