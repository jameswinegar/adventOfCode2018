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

	file, err := os.Open("Day 1/input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		j, err := strconv.Atoi(scanner.Text())
		check(err)
		i = i + j
	}
	fmt.Println(i)
}
