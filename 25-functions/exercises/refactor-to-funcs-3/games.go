package games

import (
	"encoding/json"
	"fmt"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type newJson struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Genre string `json: "genre"`
	Price int    `json: "price"`
}

func main() {

}

func newGame(id, price int, name, genre string) game {
	newGame := game{item: item{id: id, name: name, price: price}, genre: genre}
	return newGame
}

func addGame(games []game, game game) []game {
	return append(games, game)
}

func indexByID(games []game) map[int]game {
	var newMap map[int]game

	for _, v := range games {
		newMap[v.id] = v
	}
	return newMap
}

func printGame(game game) {
	fmt.Println(game)
}

func load() (games []game) {
	decoded := []newJson{}
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		panic(err)
	}
	for _, v := range decoded {
		games = addGame(games, newGame(v.ID, v.Price, v.Name, v.genre)
	}
	return games
}
