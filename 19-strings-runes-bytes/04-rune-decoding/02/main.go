package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// word := []byte("öykü")
	// fmt.Printf("%s = % [1]x \n", word)

	// for _, v := range word {
	// 	fmt.Printf("%v", strings.ToUpper(string(v)))
	// }

	// // var size int
	// _, size := utf8.DecodeRune(word)

	// copy(word[:size], bytes.ToUpper(word[:size]))
	// fmt.Printf("%s = % [1]x \n", word)
	// fmt.Println(string(103), string(111))

	word := "gökyüzü"
	bword := []byte(word)

	// ö => 2 bytes
	// ü => 2 bytes
	fmt.Println(utf8.RuneCount(bword), len(word), len(string(word[1])))
}
