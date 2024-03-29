package main

import (
	"fmt"
	"github.com/arythmetics/advent-of-code/utils"
)

var FILE_INPUT string = "input/test-1.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	r, e := utils.ParseInput(FILE_INPUT)
	check(e)
	for _, line := range r {
		fmt.Println(line)
	}
}