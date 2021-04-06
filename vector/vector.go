package main

import (
	"fmt"
)

func main() {
	fmt.Println(linespace(1, 5, 5)) // [1 2 3 4 5]

	v := []float64{1, 2, 3}
	fmt.Println("dot:", dot(v, v))
}

func dot(v1, v2 []float64) float64 {
	total := 0.0
	// for i := 0; i < len(v1); i++ {
	// 	total += v1[i] * v2[i]
	// }
	// for i := range v1 {
	// 	total += v1[i] * v2[i]
	// }
	for i, v := range v1 {
		total += v * v2[i]
	}
	return total
}

func linespace(start, stop float64, count int) []float64 {
	step := (stop - start) / float64(count-1)
	var vec []float64

	for i := 0; i < count; i++ {
		n := start + step*float64(i)
		vec = append(vec, n)
	}

	return vec
}

// How does append work?
func appendInt(s []int, n int) []int {
	idx := len(s)
	if len(s) < cap(s) { // there is room in the underlying array. we don't need to allocate more memory.
		s = s[:len(s)+1]
	} else { // no space in underlying array, need to re-allocate memory
		s1 := make([]int, 2*len(s)) // doubling the size of the array
		copy(s1, s)
		s = s1[:len(s)+1]
	}
	s[idx] = n
	return s
}