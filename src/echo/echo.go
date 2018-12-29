package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(echoOne(os.Args))

	fmt.Println(echoTwo(os.Args))

	fmt.Println(echoThree(os.Args))

}

func echoOne(input []string) string {
	//1
	var s, sep string
	for i := 1; i < len(input); i++ {
		s += sep + input[i]
		sep = " "
	}
	return s
}

func echoTwo(input []string) string {
	//2
	var q, res string
	for _, arg := range input[1:] {
		q += res + arg
		res = " "
	}
	return q
}

func echoThree(input []string) string {
	//3
	return strings.Join(input[1:], " ")
}
