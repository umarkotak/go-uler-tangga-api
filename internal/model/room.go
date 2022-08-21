package model

type (
	Room struct {
		ID                 string            `json:"id"`
		MapConfig          MapConfig         `json:"map_config"`
		PlayerMap          map[string]Player `json:"player_map"`
		PlayerRoomIndexMap map[string]Player `json:"player_room_index_map"`
		PlayerCount        int64             `json:"player_count"`
		ActivePlayer       Player            `json:"active_player"`
	}

	MapConfig struct {
		MinNumber   int64                  `json:"min_number"`
		MaxNumber   int64                  `json:"max_number"`
		Size        int64                  `json:"size"`
		Numbering   []int64                `json:"numbering"`
		Direction   []int64                `json:"direction"`
		FieldEffect map[string]FieldEffect `json:"field_effect"`
	}

	FieldEffect struct {
		FieldNumber       int64            `json:"field_number"`
		FieldNumberString string           `json:"field_number_string"`
		BenefitType       string           `json:"benefit_type"`       // [player_move]
		EffectPlayerMove  EffectPlayerMove `json:"effect_player_move"` //
	}

	EffectPlayerMove struct {
		Direction      string `json:"direction"`
		MoveCount      int64  `json:"move_count"`
		FromCoordinate int64  `json:"from_coordinate"`
		ToCoordinate   int64  `json:"to_coordinate"`
	}
)
