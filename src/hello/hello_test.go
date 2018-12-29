package main

import (
	"fmt"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	start := time.Now()
	if hello() != "Hello World !!!" {
		t.Error("Not expected value.")
	}
	fmt.Println("Time took in nano seconds.", time.Since(start).Nanoseconds())
}
