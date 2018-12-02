package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

package main

import (
"bufio"
"fmt"
"os"
"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	var id_array []string
	var twoCount, threeCount = 0,0
	var twoCounter, threeCounter = 0, 0
	var characterMap map[string]int
	characterMap = make(map[string]int)

	file, err := os.Open("Day 2/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id := scanner.Text()
		id_array = strings.Split(id, "")
		for _, character := range id_array {
			characterMap[character] = characterMap[character] + 1
		}

		twoCounter, threeCounter = 0, 0
		for _, v := range characterMap {
			if v == 2 && twoCounter == 0 {
				twoCount++
				twoCounter++
			}
			if v == 3 && threeCounter == 0 {
				threeCount++
				threeCounter++
			}
		}

		for k := range characterMap {
			delete(characterMap, k)
		}
	}
}