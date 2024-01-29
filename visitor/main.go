package main

import "fmt"

// Declare variable activeUserCount
// your code goes here


var activeUserCount int

func entry() {
	// Hint: you can use the "++" operator to increment a variable by 1
	// your code goes here

	activeUserCount +=1
	
	//fmt.Println("+1 pog")
}

func exit() {
	// Hint: you can use the "--" operator to decrement a variable by 1
	// your code goes here
	activeUserCount -=1
	//fmt.Println("-1 lul")
}

func main() {
	//fmt.Printf("We have %v visitors \n",activeUserCount)
	entry()
	
	//fmt.Printf("We have %v visitors \n",activeUserCount)
	entry()
	
	//fmt.Printf("We have %v visitors \n",activeUserCount)
	exit()
	
	//fmt.Printf("We have %v visitors \n",activeUserCount)
	exit()
	
	//fmt.Printf("We have %v visitors \n",activeUserCount)
	entry()
	
	//fmt.Printf("We have %v visitors \n",activeUserCount)
	entry()
	fmt.Println(activeUserCount)
}
