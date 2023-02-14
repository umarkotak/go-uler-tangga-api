package model

import (
	"math/rand"
	"time"
)

type (
	Room struct {
		ID                 string            `json:"id"`
		MapConfig          MapConfig         `json:"map_config"`
		PlayerMap          map[string]Player `json:"player_map"`
		PlayerRoomIndexMap map[string]Player `json:"player_room_index_map"`
		PlayerCount        int64             `json:"player_count"`
		ActivePlayer       Player            `json:"active_player"`
		MoveLogs           []MoveLog         `json:"move_logs"`
		Winners            []string          `json:"winners"`
	}

	MapConfig struct {
		Title       string                 `json:"title"`
		MinNumber   int64                  `json:"min_number"`
		MaxNumber   int64                  `json:"max_number"`
		DiceNumbers []int64                `json:"dice_numbers"`
		Size        int64                  `json:"size"`
		Numbering   []int64                `json:"numbering"`
		Direction   []int64                `json:"direction"`
		FieldEffect map[string]FieldEffect `json:"field_effect"`
		MaxPlayer   int64                  `json:"max_player"`
	}

	FieldEffect struct {
		FieldNumber           int64                  `json:"field_number"`
		FieldNumberString     string                 `json:"field_number_string"`
		BenefitType           string                 `json:"benefit_type"`            // [player_move, consumable_item]
		EffectPlayerMove      EffectPlayerMove       `json:"effect_player_move"`      //
		EffectConsumableItems []EffectConsumableItem `json:"effect_consumable_items"` // get random item from the array
	}

	EffectPlayerMove struct {
		Direction      string `json:"direction"`
		MoveCount      int64  `json:"move_count"`
		FromCoordinate int64  `json:"from_coordinate"`
		ToCoordinate   int64  `json:"to_coordinate"`
	}

	EffectConsumableItem struct {
		Name         string          `json:"name"`
		Description  string          `json:"description"`
		UsePeriodMap map[string]bool `json:"use_period"`  // [before_rolling_number, before_moving, before_end_turn]
		EffectType   string          `json:"effect_type"` // [move_n_step, teleport_to]
		Target       string          `json:"target"`      // [self, other, all]
		Value        int64           `json:"value"`
	}

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

func (r *Room) GetRandomNumber() int64 {
	randomIndex := rand.Intn(len(r.MapConfig.DiceNumbers))
	return r.MapConfig.DiceNumbers[randomIndex]
}
