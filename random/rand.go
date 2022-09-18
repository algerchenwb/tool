package random

import (
	"math/rand"
	"time"
)

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Ints(min int, max int, length int) []int {
	if min > max || length == 0 {
		return []int{}
	}
	var rt = make([]int, 0, length)
	for i := 0; i < length; i++ {
		rt = append(rt, Rand.Intn(max-min)+min)
	}
	return rt
}

func Float64s(length int) []float64 {
	var rt = make([]float64, 0, length)
	for i := 0; i < length; i++ {
		rt = append(rt, Rand.Float64())
	}
	return rt
}
