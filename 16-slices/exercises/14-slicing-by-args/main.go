package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// ---------------------------------------------------------
	// EXERCISE: Slicing by arguments
	//
	//   We've a []string, you will get arguments from the command-line,
	//   then you will print only the elements that are requested.

	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	from, to := 0, len(ships)

	switch poss := os.Args[1:]; len(poss) {
	default:
		fallthrough
	case 0:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	case 2:
		to, _ = strconv.Atoi(poss[1])
		fallthrough
	case 1:
		from, _ = strconv.Atoi(poss[0])
	}

	if l := len(ships); from < 0 || from > l || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[from:to])
}
