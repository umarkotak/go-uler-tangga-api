package model

import "time"

type (
	Room struct {
		ID                 string            `json:"id"`
		MapConfig          MapConfig         `json:"map_config"`
		PlayerMap          map[string]Player `json:"player_map"`
		PlayerRoomIndexMap map[string]Player `json:"player_room_index_map"`
		PlayerCount        int64             `json:"player_count"`
		ActivePlayer       Player            `json:"active_player"`
		MoveLogs           []MoveLog         `json:"move_logs"`
	}

	MapConfig struct {
		Title       string                 `json:"title"`
		MinNumber   int64                  `json:"min_number"`
		MaxNumber   int64                  `json:"max_number"`
		Size        int64                  `json:"size"`
		Numbering   []int64                `json:"numbering"`
		Direction   []int64                `json:"direction"`
		FieldEffect map[string]FieldEffect `json:"field_effect"`
		MaxPlayer   int64                  `json:"max_player"`
	}

	FieldEffect struct {
		FieldNumber       int64            `json:"field_number"`
		FieldNumberString string           `json:"field_number_string"`
		BenefitType       string           `json:"benefit_type"`       // [player_move, boost_item]
		EffectPlayerMove  EffectPlayerMove `json:"effect_player_move"` //
		EffectBoostItem   EffectBoostItem  `json:"effect_boost_item"`
	}

	EffectPlayerMove struct {
		Direction      string `json:"direction"`
		MoveCount      int64  `json:"move_count"`
		FromCoordinate int64  `json:"from_coordinate"`
		ToCoordinate   int64  `json:"to_coordinate"`
	}

	EffectBoostItem struct{}

	MoveLog struct {
		Log       string    `json:"log"`
		Timestamp time.Time `json:"timestamp"`
	}
)

func (r *Room) WriteMoveLog(log string) {
	r.MoveLogs = append(r.MoveLogs, MoveLog{
		Log:       log,
		Timestamp: time.Now(),
	})
}
