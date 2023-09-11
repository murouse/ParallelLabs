package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
)

// go run .\lab01\task02\ 18446744073709551615

func main() {
	arg := os.Args[1]
	fmt.Printf("original argument %v with type %v\n", arg, reflect.TypeOf(arg))

	res1, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("parsed argument %v with type %v\n", res1, reflect.TypeOf(res1))

	res2, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("parsed argument %v with type %v\n", res2, reflect.TypeOf(res2))

	res3_, err := strconv.ParseUint(arg, 10, 32)
	if err != nil {
		log.Panic(err)
	}
	res3 := uint32(res3_)
	fmt.Printf("parsed argument %v with type %v\n", res3, reflect.TypeOf(res3))

	res4_, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		log.Panic(err)
	}
	res4 := int32(res4_)
	fmt.Printf("parsed argument %v with type %v\n", res4, reflect.TypeOf(res4))

	fmt.Printf("math.MaxUint64: %v\n", uint64(math.MaxUint64))
	fmt.Printf("math.MaxInt64: %v\n", math.MaxInt64)
	fmt.Printf("math.MaxUint32: %v\n", math.MaxUint32)
	fmt.Printf("math.MaxInt32: %v\n", math.MaxInt32)

	fmt.Printf("math.MinUint64: %v\n", 0)
	fmt.Printf("math.MinInt64: %v\n", math.MinInt64)
	fmt.Printf("math.MinUint32: %v\n", 0)
	fmt.Printf("math.MinInt32: %v\n", math.MinInt32)

}
