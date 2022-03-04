package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"
	str := strings.Split(data, " ")

	//   1. Convert the string to an []int
	newNums := []int{}

	for _, v := range str {
		intVar, _ := strconv.Atoi(v)
		newNums = append(newNums, intVar)
	}
	//   2. Print the slice
	fmt.Println(newNums)

	//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
	//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
	even := newNums[:3]
	odd := newNums[3:]

	fmt.Println(even)
	fmt.Println(odd)

	//   5. Slice it for the two numbers at the middle
	fmt.Println(newNums[2:4])

	//   6. Slice it for the first two numbers
	fmt.Println(newNums[:2])

	//   7. Slice it for the last two numbers (use the len function)
	fmt.Println(newNums[len(newNums)-2:])

	//   8. Slice the evens slice for the last one number
	fmt.Println(even[2:])

	//   9. Slice the odds slice for the last two numbers
	fmt.Println(odd[1:])
}
