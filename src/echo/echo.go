package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//2
	var q, res string
	for _, arg := range os.Args[1:] {
		q += res + arg
		res = " "
	}
	fmt.Println(q)

	//3
	fmt.Println(strings.Join(os.Args[1:], " "))
}
