// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	newData := strings.Split(data, "\n")
	fmt.Println(newData)

	var (
		locations []string
		sizes     []int
		beds      []int
		baths     []int
		prices    []int
	)
	for _, v := range newData {

		cols := strings.Split(v, separator)

		locations = append(locations, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range newData {
		fmt.Printf("%-15s %-15d %-15d %-15d %-15d\n", locations[i], sizes[i], beds[i], baths[i], prices[i])
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	fmt.Printf("%-15s %-15.2f %-15.2f %-15.2f %-15.2f\n", "",
		float64(sum(sizes))/float64(len(sizes)),
		float64(sum(beds))/float64(len(beds)),
		float64(sum(baths))/float64(len(baths)),
		float64(sum(prices))/float64(len(prices)))
}

func sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
