package main

import "fmt"

func main() {
	n := average2(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average2(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%t \n", sf)
	var total float64
	for _, v := range sf  {
		total += v
	}
	return total / float64(len(sf))
}