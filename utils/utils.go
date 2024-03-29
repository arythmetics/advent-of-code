package utils

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseInput(path string) ([]string, error) {
	var lines []string
	f, e := os.Open(path)
	check(e)
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines, fileScanner.Err()
}