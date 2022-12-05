package main

import "fmt"

func main() {
	//var a [3]string
	a := make([]string, 0)
	counter := 0

	for counter < 3 {
		a = append(a, "x")
		fmt.Print(len(a))
		counter++
	}

}
