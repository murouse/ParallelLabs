package main

import (
	"ParallelLabs/lab01/task05/utils"
	"fmt"
)

func main() {
	elems := []int{1, 3, -3, 0, -2, 3, 0, 2, -6, 0, 1}
	fmt.Println(elems)

	predicates := utils.Predicates[int]{
		func(x *int) bool { return *x == 0 },
		func(x *int) bool { return *x < 0 && *x%2 == 0 },
		func(x *int) bool { return *x > 0 && *x%2 == 1 },
	}

	for _, predicate := range predicates {
		elems = utils.RemoveIf[int](elems, predicate)
		fmt.Println(elems)
	}
}
