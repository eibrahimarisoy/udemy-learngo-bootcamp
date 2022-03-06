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
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {
	var (
		words map[string]bool
		count int
	)
	words = make(map[string]bool)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	re, err := regexp.Compile(`[^A-Za-z]+`)
	if err != nil {
		fmt.Println(err)
		return
	}

	for in.Scan() {
		text := re.ReplaceAllString(in.Text(), "")
		text = strings.ToLower(text)
		if _, ok := words[text]; !ok {
			words[text] = true
		}
		count++
	}

	fmt.Println(len(words))
	fmt.Println(count)

}
