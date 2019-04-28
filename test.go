package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "1234"
	size := len(str)

	for _, r := range str {
		// fmt.Println(r)
		i, _ = strconv.Atoi(r)
		fmt.Println(i)
	}
}
