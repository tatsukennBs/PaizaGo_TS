//go:build ignore

package main

import (
	"fmt"
)

func main() {

	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	s = append(s, "e")
	s = append(s, "f", "g", "h")
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	c := make([]string, 7)
	copy(c, s)
	fmt.Println(c)

	slice := s[2:5]
	fmt.Println(slice)

	slice = s[:5]
	fmt.Println(slice)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
