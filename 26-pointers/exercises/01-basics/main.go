// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var blank *computer
	fmt.Printf("%T", blank) //*main.computer
	fmt.Println()
	fmt.Println(blank) //nil

	// compare the pointer variable to a nil value, and say it's nil
	if blank == nil {
		fmt.Println("it's nil")
	}

	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "Apple"}
	fmt.Printf("%T", apple)

	// put the apple into a new pointer variable
	newApple := apple

	// compare the apples: if they are equal say so and print their addresses
	if apple == newApple {
		fmt.Println("they are equal")
		fmt.Printf("p, %p \n", apple)
		fmt.Printf("t, %p", newApple)
	} else {
		fmt.Println("they are not equal")
	}

	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{"sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if sony != apple {
		fmt.Println()
		fmt.Println("They are not equal")
		fmt.Printf("sony %#p \n", sony)
		fmt.Printf("apple %#p \n", apple)
	}
	// put apple's value into a new ordinary variable
	lastApple := *apple
	fmt.Println(lastApple)

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple pointer %p %p \n", &apple, apple)
	fmt.Printf("last apple pointer %d \n", &newApple)

	// // compare the value that is pointed by the apple and the new variable
	// // if they are equal say so

	if *apple == lastApple {
		fmt.Println("they are equal")
		// print the values:

		// the value that is pointed by the apple and the new variable
		fmt.Printf("apple      : %+v - lastApple %+v \n", *apple, lastApple)
	}

	// create a new function: change
	// the func can change the given computer's brand to another brand
	// // change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	fmt.Println("changed sony", *sony)

	// // write a func that returns the value that is pointed by the given *computer
	// // print the returned value
	fmt.Printf("appleVal  : %+v\n", valueOf(apple))

	// // write a new constructor func that returns a pointer to a computer
	// // and call the func 3 times and print the returned values' addresses
	fmt.Printf("%p\n", newComputer("nokia"))
}

func change(c *computer, brand string) {
	c.brand = brand
}

func valueOf(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	new := &computer{brand: brand}
	return new
}
