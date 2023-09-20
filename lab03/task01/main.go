package main

import (
	"ParallelLabs/lab03/task01/utils"
	"fmt"
	"math/rand"
)

func main() {
	const n = 3

	fmt.Println(utils.Kiss(n, [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}))
}
