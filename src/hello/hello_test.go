package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if hello() != "Hello World !!!" {
		t.Error("Not expected value.")
	}
}
