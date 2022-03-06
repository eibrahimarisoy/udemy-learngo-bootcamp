package games

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
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

func load() []game {
	games := []game{}
	games = addGame(games, newGame(0, 5, "call of", "war"))
	games = addGame(games, newGame(0, 5, "gta", "fun"))
	return games
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
