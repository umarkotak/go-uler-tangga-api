package main

import (
	"encoding/json"
	"fmt"
)

type World struct {
	RoomMap map[string]Room
}

type Room struct {
	State string
}

func main() {
	world := World{
		RoomMap: map[string]Room{},
	}

	room := Room{
		State: "init",
	}

	world.RoomMap["room1"] = room

	res, _ := json.Marshal(world)
	fmt.Println("ONE", string(res))

	tempRoom := world.RoomMap["room1"]
	tempRoom.State = "HELLO THERE"

	res, _ = json.Marshal(world)
	fmt.Println("TWO", string(res))
}
