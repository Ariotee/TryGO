package main

import (
	"fmt"
)

func main() {
	var first, second int = 1, 2

	var sum, multiplay int = SumAndMultiplay(first, second)
	fmt.Println(sum, multiplay)
}

func SumAndMultiplay(first, second int) (int, int) {
	var sum, multiplay int = first + second, first * second
	return sum, multiplay
}
