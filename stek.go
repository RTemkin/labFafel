package main

import (
	"fmt"
	"strings"
)

func main() {

	stack := []string{}

	a := "world"

	stack = append(stack, a)

	//b := "Hello"

	stack2 := []int{0, 1, 2, 3}

	//stack2 = stack2[:len(stack2)-1]
	stack2 = stack2[1:3]

	fmt.Println(stack2)

	line := "a or ( b and r or u )"
	tokens := strings.Fields(line)

	fmt.Println(tokens)

}
