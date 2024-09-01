package quests

func Swap(x *int, y *int) (*int, *int) {
	*x, *y = *y, *x
	return x, y
}
