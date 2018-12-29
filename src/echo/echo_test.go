package main

import (
	"testing"
)

func TestEchoOne(t *testing.T) {
	s := []string{"", "This", "is", "to", "Test"}
	if echoOne(s) != "This is to Test" {
		t.Error("Expectation failed.")
	}
}

func TestEchoTwo(t *testing.T) {
	s := []string{"", "This", "is", "to", "Test"}
	if echoTwo(s) != "This is to Test" {
		t.Error("Expectation failed.")
	}
}

func TestEchoThree(t *testing.T) {
	s := []string{"", "This", "is", "to", "Test"}
	if echoThree(s) != "This is to Test" {
		t.Error("Expectation failed.")
	}
}
