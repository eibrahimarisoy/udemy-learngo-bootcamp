package main

import (
	"bytes"
	"fmt"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 1. yol
	// for _, v := range png {
	// 	header = append(header, v)
	// }

	// 2. yol
	header = append(header, png...)

	if bytes.Equal(header, png) {
		fmt.Println("They are equal.")
	}
}
