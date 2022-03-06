// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	floatOne, floatTwo := 10.1, 15.2
	fmt.Println(floatOne, floatTwo)
	swap(&floatOne, &floatTwo)
	fmt.Println(floatOne, floatTwo)

	// 2.nd section
	a, b := &floatOne, &floatTwo
	// var a, b *float64
	// a, b = &floatOne, &floatTwo
	a, b = swapAddress(a, b)
	fmt.Println(a, b)

}

func swap(one, two *float64) {
	fmt.Println(one, two)
	*one, *two = *two, *one
	fmt.Printf("%p    %p\n", one, two)
}

func swapAddress(one, two *float64) (*float64, *float64) {
	return two, one
}
