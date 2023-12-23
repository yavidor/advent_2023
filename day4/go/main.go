package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

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
func main() {
	fileData, err := fs.ReadFile(os.DirFS(".."), "example.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines [][][]string = makeLines(fileData)
	fmt.Println(lines)
}
