package main

import "fmt"

func calculate(a int, b int) (results []float64) {
  // your code goes here
	s := a + b
	d := a - b
	p := a * b
	q := a / b

	results = append(results, float64(s), float64(d), float64(p), float64(q))

	//results = append(results, float64(a+b), float64(a-b), float64(a*b), float64(a/b))


	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}