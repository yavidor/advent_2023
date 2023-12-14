package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	value   int
	digits  []rune
	incides [][]int
}

func main() {
	file_data, err := fs.ReadFile(os.DirFS(".."), "input.txt")
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	// 	file_data :=
	// 		`467..114..
	// ...*......
	// ..35..633.
	// ......#...
	// 617*......
	// .....+.58.
	// ..592.....
	// ......755.
	// ...$.*....
	// .664.598..`
	var lines [][]rune = [][]rune{}
	var numbers []Number
	newNumber := true
	for _, line := range strings.Split(string(file_data), "\n") {
		lines = append(lines, []rune(line))
	}
	lines = lines[:len(lines)-1]
	for row, i := range lines {
		for col, j := range i {
			if unicode.IsDigit(j) {
				if newNumber {
					numbers = append(numbers, Number{value: 0, digits: []rune{j}, incides: [][]int{{row, col}}})
					newNumber = false
				} else {
					numbers[len(numbers)-1].digits = append(numbers[len(numbers)-1].digits, j)
				}
				if row > 0 {
					numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row - 1, col})
					if col > 0 {
						numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row - 1, col - 1})
					}
					if col < len(lines[0])-1 {
						numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row - 1, col + 1})
					}
				}
				if row < len(lines)-1 {
					numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row + 1, col})
					if col > 0 {
						numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row + 1, col - 1})
					}
					if col < len(lines[0])-1 {
						numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row + 1, col + 1})
					}
				}
				if col > 0 {
					numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row, col - 1})
				}
				if col < len(lines[0])-1 {
					numbers[len(numbers)-1].incides = append(numbers[len(numbers)-1].incides, []int{row, col + 1})
				}
			} else {
				if !newNumber {
					numbers[len(numbers)-1].value, _ = strconv.Atoi(string(numbers[len(numbers)-1].digits))
				}
				newNumber = true
			}
		}
	}
	for _, number := range numbers {
		isPart := false
		for _, rowIndex := range number.incides {
			if !(unicode.IsDigit(lines[rowIndex[0]][rowIndex[1]]) || lines[rowIndex[0]][rowIndex[1]] == '.') {
				isPart = true
				break
			}
		}
		if isPart {
			sum += number.value
		}
	}
	fmt.Println(sum)
}
