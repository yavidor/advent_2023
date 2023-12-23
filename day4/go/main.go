package main

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"reflect"
	"strings"
)

func contains[T any](s []T, v T) bool {
	for _, vS := range s {
		if reflect.DeepEqual(vS, v) {
			return true
		}
	}
	return false
}
func makeLines(data []byte) [][][]string {
	inputLines := strings.Split(string(data), "\n")
	outputLines := make([][][]string, 0, len(inputLines))
	for _, line := range inputLines {
		lineContent := line[strings.Index(line, ":")+2:]
		fields := strings.Split(lineContent, "|")
		var subLine [][]string
		for _, field := range fields {
			subLine = append(subLine, strings.Fields(field))
		}
		outputLines = append(outputLines, subLine)

	}
	return outputLines
}
func calcCardValue(card [][]string) int {
	ret := 0
	for _, v := range card[1] {
		if contains(card[0], v) {
			ret += 1
		}
	}
	return ret
}
func partOne(lines [][][]string) int {
	sum := 0
	for _, v := range lines {
		sum += int(math.Pow(2, float64(calcCardValue(v))-1))
	}
	return sum
}
func main() {
	fileData, err := fs.ReadFile(os.DirFS(".."), "input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines [][][]string = makeLines(fileData)
	fmt.Println(lines)
	fmt.Println(partOne(lines))
}
