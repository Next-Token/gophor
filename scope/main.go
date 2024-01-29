package main

import (
	"fmt"
	"strings"
)


func modify(y int) int {
	y += 15
	return y
}
func main() {
	y := 20
	modify(y)
	fmt.Println(y)
}