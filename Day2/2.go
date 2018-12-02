package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func hammingDistance(s1 string, s2 string) (distance int, err error) {
	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) != len(r2) {
		err = errors.New("Hamming distance of different sized strings.")
		return
	}

	for i, v := range r1 {
		if r2[i] != v {
			distance += 1
		}
	}
	return
}

func matchingCharacters(s1 string, s2 string) (result string, err error) {
	if len(s1) != len(s2) {
		err = errors.New("Cannot compare elements of  different sized strings.")
		return
	}

	for i, char := range s1 {
		if s1[i] == s2[i] {
			result += string(char)
		}
	}
	return
}

func main() {
	lines, err := readLines("Day 2/input.txt")
	check(err)

	for i, line1 := range lines {
		for _, line2 := range lines[i+1:] {
			distance, err := hammingDistance(line1, line2)
			check(err)
			if distance == 1 {
				result, err := matchingCharacters(line1, line2)
				check(err)
				fmt.Println(result)
			}
		}
	}
}