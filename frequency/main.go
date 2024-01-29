package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string)  map[string]int {
	// TODO: implement this function
	words := strings.Fields(text)
	jj := make(map[string]int)



	for _, j := range words{
		jj[j]++
		fmt.Printf("we got %v %v\n", j, words[0])
	}
	return jj
}

func main() {
	text := "The the quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}