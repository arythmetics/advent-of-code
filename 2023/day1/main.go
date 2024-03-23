package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var FILE_INPUT string = "puzzle-input-2.txt"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func castToInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func DayOnePartOne(fileName string) int {
	data, err := os.ReadFile(fileName)
	check(err)
	var list []string = strings.Split(string(data), "\n")
	var calibrationValues []int
	for _, val := range list {
		var firstNumber int
		for i := 0; i < len(val); i++ {
			letter := rune(val[i])
			if unicode.IsDigit(letter) {
				firstNumber = castToInt(string(letter))
				break
			}
		}
		var secondNumber int
		for i := len(val)-1; i >= 0; i-- {
			letter := rune(val[i])
			if unicode.IsDigit(letter) {
				secondNumber = castToInt(string(letter))
				break
			}
		}
		fullNumber := firstNumber*10 + secondNumber
		calibrationValues = append(calibrationValues, fullNumber)
	}
	sumCalibrationValues := sum(calibrationValues)
	return sumCalibrationValues
}

func createLetterMap(numericDigits [10]string) map[string]string {
	letterMap := make(map[string]string)
	letterKeys := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
	for i, letter := range(letterKeys) {
		letterMap[letter] = numericDigits[i]
	}
	return letterMap
}

var NUMERIC_DIGITS = [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var LETTER_MAP = createLetterMap(NUMERIC_DIGITS)

func findDigit(s string) string {
	for _, val := range NUMERIC_DIGITS {
		if strings.Contains(s, val) {
			digit := val
			return digit
		}
	}
	for val := range LETTER_MAP {
		if strings.Contains(s, val) {
			digit := LETTER_MAP[val]
			return digit
		}
	}
	return ""
}

func isEmptyString(s string) bool {
	return len(s) == 0
}

func DayOnePartTwo(fileName string) int {
	data, err := os.ReadFile(fileName)
	check(err)
	var calibrationValues []int
	var list []string = strings.Split(string(data), "\n")
	for _, val := range list {
		var firstNumber int
		for i := 0; i <= len(val); i++ {
			digit := findDigit(val[0:i])
			if !isEmptyString(digit) {
				firstNumber = castToInt(digit)
				break
			}
		}
		var secondNumber int
		for i := len(val)-1; i >= 0; i-- {
			digit := findDigit(val[i:])
			if !isEmptyString(digit) {
				secondNumber = castToInt(digit)
				break
			}
		}
		fullNumber := firstNumber*10 + secondNumber
		calibrationValues = append(calibrationValues, fullNumber)
	}
	return sum(calibrationValues)
}

func main() {
	result := DayOnePartTwo(FILE_INPUT)
	fmt.Println(result)
}