package main

import "fmt"

func main() {
	fmt.Print("Enter feet: ")
	var ft float64
	fmt.Scanf("%f", &ft)

	m := ft * 0.3048

	fmt.Println(m, "meters")
}
