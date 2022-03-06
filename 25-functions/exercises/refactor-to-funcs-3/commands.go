package main

func main() {

}

func runCmd(input string, []game, map[int]game) bool{

		// --- runCmd start ---
		fmt.Println()

		cmd := strings.Fields(input)
		if len(cmd) == 0 {
			return true
		}

		switch input[0] {
		case "quit":
			return cmdQuit()
		case "id":
			return cmdById(cmd, games, byID)
		case "list":
			return cmdList(games)
		case "save":
			return cmdSave(games)
		}
		return false

}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList([]game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}


func cmdById(cmd []string, []game, map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry. I don't have the game")
		return true
	}

	printGame(g)
	return true
}

func cmdSave(cmd []game) bool {
	type newJson struct {
		ID    int    `json: "id"`
		Name  string `json: "name"`
		Genre string `json: "genre"`
		Price int    `json: "price"`
	}
	var newDatas []newJson


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

	byteData, err := json.MarshalIndent(newSlices, "", "\t")
	if err != nil {
		panic(err)
		return true
	}
	fmt.Println(string(byteData))
	return false
}