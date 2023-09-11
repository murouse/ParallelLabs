package main

import (
	"log"
	"math"
	"testing"
)

func TestFuncNormal(t *testing.T) {
	res := NormalPDF(5, 5.25, math.Sqrt(0.35))
	log.Printf("%.8f", res)
	//0.61674
}
