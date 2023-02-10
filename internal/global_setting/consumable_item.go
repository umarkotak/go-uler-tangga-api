package global_setting

import "github.com/umarkotak/go-uler-tangga-api/internal/model"

var (
	EffectConsumableItemsSet1 = []model.EffectConsumableItem{CONSUMABLE_MAJU_5_LANGKAH}

	CONSUMABLE_MAJU_5_LANGKAH = model.EffectConsumableItem{
		Name:       "maju 5 langkah",
		EffectType: "move_n_step",
		Target:     "self",
		Value:      5,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}
)
