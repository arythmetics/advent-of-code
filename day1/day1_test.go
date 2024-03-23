package main

import (
	"testing"
)

func TestDayOne1(t *testing.T) {
	res := DayOnePartOne("test-input-1.txt")
	want := 142
	if res != want {
		t.Fatalf(`DayOne1("test-input-1.txt") = %d, want match for %d`, res, want)
	}
}

func TestDayOne2(t *testing.T) {
	res := DayOnePartTwo("test-input-2.txt")
	want := 314
	if res != want {
		t.Fatalf(`DayOne2("test-input-2.txt) = %d, want match for %d`, res, want)
	}
}