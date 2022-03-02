package main

import (
	"fmt"
	"time"
)

func main() {
	var pizza []string
	var departures []time.Time
	var graduations []int
	var lights []bool
	pizza = append(pizza, "cheese")
	fmt.Println(pizza)

	departures = append(departures, time.Now())
	fmt.Println(departures)

	graduations = append(graduations, 2020)
	fmt.Println(graduations)

	lights = append(lights, true, false)
	fmt.Println(lights)
}
