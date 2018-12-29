package main

import (
	"fmt"
	"testing"
	"time"
)

func TestEchoOne(t *testing.T) {
	start := time.Now()
	s := []string{"", "This", "is", "to", "Test"}
	if echoOne(s) != "This is to Test" {
		t.Error("Expectation failed.")
	}
	fmt.Println("Time took in nano seconds.", time.Since(start).Nanoseconds())
}

func TestEchoTwo(t *testing.T) {
	start := time.Now()
	s := []string{"", "This", "is", "to", "Test"}
	if echoTwo(s) != "This is to Test" {
		t.Error("Expectation failed.")
	}
	fmt.Println("Time took in nano seconds.", time.Since(start).Nanoseconds())
}

func TestEchoThree(t *testing.T) {
	start := time.Now()
	s := []string{"", "This", "is", "to", "Test"}
	if echoThree(s) != "This is to Test" {
		t.Error("Expectation failed.")
	}
	fmt.Println("Time took in nano seconds.", time.Since(start).Nanoseconds())
}
