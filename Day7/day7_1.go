package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	dirTree := make(map[string]int)
	newDirTree := make(map[string]int)
	currentDir := ""
	lineNbr := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//figure out command or output
		if string(scanner.Text()[0]) == "$" {
			// Command

			if strings.Split(scanner.Text(), " ")[1] == "cd" {

				if strings.Split(scanner.Text(), " ")[2] == "/" {

					currentDir = "/"

				} else if strings.Split(scanner.Text(), " ")[2] == ".." {

					//new shit
					fmt.Print("Move Up ..\n\n\n")
					currentDirSlice := strings.Split(currentDir, "/")
					//currentDirSlice := currentDirSlice[0 : len(currentDir)-1]
					newstring := ""
					for x := 0; x < len(currentDirSlice)-1; x++ {
						newstring += string(currentDirSlice[x]) + "/"
					}
					currentDir = newstring[0 : len(newstring)-1]

				} else {
					if currentDir == "/" {
						currentDir += strings.Split(scanner.Text(), " ")[2]
					} else {
						currentDir += "/" + strings.Split(scanner.Text(), " ")[2]
					}

				}
			}

		} else if string(scanner.Text()[0]) == "d" {
			// Directory

			dir := strings.Split(scanner.Text(), " ")[1]

			if currentDir == "/" {
				dirTree[dir] = 0
			} else if _, ok := dirTree[currentDir+dir]; !ok {
				dirTree[currentDir+"/"+dir] = 0
			}

		} else if _, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[0]); err == nil {
			// File size

			sizeStr := strings.Split(scanner.Text(), " ")[0]
			sizeInt, err := strconv.Atoi(sizeStr)
			check(err)
			dirTree[currentDir] += sizeInt

		}

		fmt.Printf("Line: %d\tCurrent dir: %s\n", lineNbr, currentDir)
		lineNbr += 1

	}

	fmt.Print(dirTree)

	for dir, size := range dirTree {
		newDirTree[dir] = size
	}

	for ndir, _ := range newDirTree {
		for dir, size := range dirTree {
			if strings.HasPrefix(dir, ndir) && ndir != dir {
				newDirTree[ndir] += size
			}
		}
	}

	fmt.Printf("\n\n")
	fmt.Print(newDirTree)

	totalSize := 0

	for _, size := range newDirTree {
		if size < 100000 {
			totalSize += size
		}
	}

	fmt.Printf("\n\nTotal Size: %d\n", totalSize)
}
