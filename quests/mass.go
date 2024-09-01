package quests

import "fmt"

func Min(x []int) int {
	g := x[0]
	for i := 0; i < len(x); i++ {
		if g > x[i] {
			g = x[i]
		}
	}
	return g
}

func MinComp() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(Min(x))
}
