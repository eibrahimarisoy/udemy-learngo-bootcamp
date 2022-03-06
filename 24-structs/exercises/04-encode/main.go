// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

// use your solution from the previous exercise
func main() {

	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := map[int]game{
		1: {
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		2: {
			item:  item{id: 2, name: "x-com 2", price: 30},
			genre: "strategy",
		},
		3: {
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
  > quit   : quits
`)

		if !in.Scan() {
			break
		}
		fmt.Printf("%T", in.Text())

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue

			}

			v, ok := games[id]

			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue

			}
			fmt.Println(v)

		case "list":
			for _, v := range games {

				fmt.Printf("%d. %s (%s) is $%v\n", v.id, v.name, v.genre, v.price)
			}

		case "quit":
			fmt.Println("good bye")
			return

		case "save":
			type newJson struct {
				ID    int    `json: "id"`
				Name  string `json: "name"`
				Genre string `json: "genre"`
				Price int    `json: "price"`
			}
			newSlices := []newJson{}

			fmt.Println("length", len(games))

			for _, v := range games {
				new := newJson{
					ID:    v.item.id,
					Name:  v.item.name,
					Genre: v.genre,
					Price: v.item.price,
				}
				newSlices = append(newSlices, new)
				// fmt.Printf("%T", v)
			}
			fmt.Println(newSlices)

			byteData, err := json.MarshalIndent(newSlices, "", "\t")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(byteData))
			return
		}

	}
}
