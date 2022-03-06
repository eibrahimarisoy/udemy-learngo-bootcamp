// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

func main() {
	words := make(map[string]bool)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		if _, ok := words[strings.ToLower(in.Text())]; !ok {
			words[strings.ToLower(in.Text())] = true
		} else {
			fmt.Println("TWICE")
			return
		}
	}
}
