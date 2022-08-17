package singleton

import "github.com/umarkotak/go-uler-tangga-api/internal/model"

var (
	world *model.World
)

func Initialize() {
	world = &model.World{
		RoomMap: map[string]model.Room{},
	}
}

func GetWorld() *model.World {
	return world
}
