package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isOverMax(gameBalls map[string][]int) bool {
	maxBalls := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for color, balls := range gameBalls {
		for _, ball := range balls {
			if ball > maxBalls[color] {
				return false
			}
		}
	}
	return true
}

func main() {
	sum := 0
	var reBalls = regexp.MustCompile(`(?m)(\d{1,2} \w{3,5})`)
	var reNums = regexp.MustCompile(`(?m)\d+`)
	var reChars = regexp.MustCompile(`(?m)[a-zA-Z]+`)
	file_data, err := fs.ReadFile(os.DirFS("."), "input.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var gamesArr []string = strings.Split(string(file_data), "\n")
	for gameIndex, gameValue := range gamesArr {
		fmt.Println("Game ", gameIndex+1)
		gameBalls := map[string][]int{
			"red":   {},
			"green": {},
			"blue":  {},
		}
		for _, match := range reBalls.FindAllString(gameValue, -1) {
			matchNumber, _ := strconv.Atoi(reNums.FindAllString(match, -1)[0])
			matchColor := reChars.FindAllString(match, -1)[0]
			gameBalls[matchColor] = append(gameBalls[matchColor], matchNumber)
		}
		fmt.Println(gameBalls)
		isValid := isOverMax(gameBalls)
		if isValid {
			sum += gameIndex + 1
		}

	}
	fmt.Println(sum)
}
