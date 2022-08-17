package model

type (
	World struct {
		RoomMap map[string]Room `json:"room_map"`
	}
)
