// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

// func main() {
// 	words := []string{
// 		"cool",
// 		"güzel",
// 		"jīntiān",
// 		"今天",
// 		"read 🤓",
// 	}

// 	// Hint: Use len and utf8.RuneCountInString
// 	for _, word := range words {
// 		fmt.Printf("--- %s  %q---\n", word, word)

// 		// Print the byte and rune length of the strings
// 		fmt.Println("\tbyte ", len(word), "rune ", utf8.RuneCountInString(word))

// 		// Print the bytes of the strings in hexadecimal
// 		fmt.Printf("\tbytes: % x", word)

// 		// Print the runes of the strings in hexadecimal
// 		fmt.Print("\trunes   :")
// 		for _, w := range words {
// 			fmt.Printf("% x ", w)
// 		}
// 		fmt.Println()

// 		// Print the runes of the strings as rune literals
// 		// Hint: Use for range
// 		fmt.Print("\trune literals   :")

// 		for _, w := range word {
// 			fmt.Printf("%q", w)
// 		}
// 		fmt.Println()

// 		// Print the first rune and its byte size of the strings
// 		// Hint: Use utf8.DecodeRuneInString

// 		r, i := utf8.DecodeRuneInString(word)
// 		fmt.Printf("\tfirst  : %q (%d bytes)", r, i)

// 		// Print the last rune of the strings
// 		// Hint: Use utf8.DecodeLastRuneInString
// 		r, i = utf8.DecodeLastRuneInString(word)
// 		fmt.Printf("\tlast  : %q (%d bytes)", r, i)

// 		// Slice and print the first two runes of the strings
// 		_, first := utf8.DecodeRuneInString(word)
// 		_, second := utf8.DecodeRuneInString(word[first:])
// 		fmt.Printf("\tfirst 2 : %q\n", word[:first+second])

// 		// Slice and print the last two runes of the strings
// 		_, last1 := utf8.DecodeLastRuneInString(word)
// 		_, last2 := utf8.DecodeLastRuneInString(word[:len(word)-last1])
// 		fmt.Println(last1, last2)
// 		fmt.Printf("\tlast 2 : %q\n", word[len(word)-last2-last1:])

// 		// Convert the string to []rune
// 		// Print the first and last two runes
// 		rs := []rune(word)
// 		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
// 		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))

// 		fmt.Println()
// 		fmt.Println(strings.Repeat("-", 45))
// 	}

// }
