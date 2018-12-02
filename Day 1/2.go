package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	var i = 0
	var m map[int]int
	m = make(map[int]int)
	m[0] = 1

	for {
		file, err := os.Open("Day 1/input.txt")
		check(err)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			j, err := strconv.Atoi(scanner.Text())
			check(err)
			i = i + j

			_, ok := m[i]
			if ok {
				fmt.Println(i)
				os.Exit(0)
			} else {
				m[i] = 1
			}
		}
		file.Close()
	}
}