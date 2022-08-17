package model

type (
	Identity struct {
		ID                    string `json:"id"`
		Name                  string `json:"name"`
		RoomID                string `json:"room_id"`
		RoomPlayerIndex       int    `json:"room_player_index"`
		RoomPlayerIndexString string `json:"room_player_index_string"`
	}
)
