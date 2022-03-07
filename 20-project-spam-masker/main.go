// #1- Get and check the input
// #2- Create a byte buffer and use it as the output
// #3- Write input to the buffer as it is and print it
// #4- Detect the link
// #5- Mask the link
// #6- Stop masking when whitespace is detected
// #7- Put a http:// prefix in front of the masked link
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please provide a string")
		return
	}

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
	)
	in := false
	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	for i := 0; i < size; i++ {

		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			fmt.Printf(`text[%d : %[1]d + %d] = `, i, nlink)
			fmt.Println(text[i : i+nlink])
			in = true
		}
		c := text[i]
		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = byte(mask)
		}
		buf = append(buf, c)
	}
	fmt.Println(string(buf))
}
