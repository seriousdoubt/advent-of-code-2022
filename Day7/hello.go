package main

import (
	"fmt"
)

func main() {

	line := make([]string, 3)

	line[0] = "dir1"

	line[1] = "dir2"

	line[2] = "dir3"
	fmt.Print("\n\n")
	fmt.Print(line)

	line = line[0 : len(line)-1]
	fmt.Print("\n\n")
	fmt.Print(line)

	newstring := "/"
	for x := 0; x < len(line); x++ {
		newstring += string(line[x]) + "/"
	}

	fmt.Printf("New Line: %s", newstring)
}
