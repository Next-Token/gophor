package main

import "fmt"

func discountedPrice(product string, price float64) (newP float64){
	var p float64
	if  product == "apples" {
		p = .1
	} else if product == "bananas"{
		p = .2
	} else {
		p = 0
	}
	disc := price * p
	newP = price -  disc
	return newP
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
