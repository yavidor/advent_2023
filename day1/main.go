package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convToDigits(line string) string {
	digitsArr := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for k := range line {
		for ak, av := range digitsArr {
			re := regexp.MustCompile(av)
			formattedLine := re.ReplaceAllString(line[0:k+1], strconv.Itoa(ak+1)) + line[k+1:]
			if formattedLine != line {
				return convToDigits(formattedLine)
			}
		}
	}
	return line
}

func main() {
	re := regexp.MustCompile(`(?m)\D`)
	file_data, err := fs.ReadFile(os.DirFS("."), "input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines []string = strings.Split(string(file_data), "\n")
	var numbers []int
	for i := 0; i < len(lines); i++ {
		lines[i] = convToDigits(lines[i])
		digits := re.ReplaceAllString(lines[i], "")
		if digits != "" {
			first := string(digits[0])
			last := string(digits[len(digits)-1])
			number, err := strconv.Atoi(first + last)
			if err != nil {
				fmt.Println(err)
			}
			numbers = append(numbers, number)
		}
	}
	var sum int
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)
}
