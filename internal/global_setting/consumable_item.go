package global_setting

import "github.com/umarkotak/go-uler-tangga-api/internal/model"

var (
	EffectConsumableItemsSet1 = []model.EffectConsumableItem{
		CONSUMABLE_MAJU_5_LANGKAH_SELF,
		CONSUMABLE_MAJU_2_LANGKAH_SELF,
		CONSUMABLE_MAJU_3_LANGKAH_SELF,
		CONSUMABLE_MAJU_1_LANGKAH_ALL,
		CONSUMABLE_MAJU_2_LANGKAH_ALL,
		CONSUMABLE_MAJU_3_LANGKAH_ALL,
		CONSUMABLE_MAJU_5_LANGKAH_ALL,
		CONSUMABLE_MUNDUR_10_LANGKAH_ALL,
		CONSUMABLE_MUNDUR_1_LANGKAH_ALL,
	}

	CONSUMABLE_MAJU_5_LANGKAH_SELF = model.EffectConsumableItem{
		Name:        "maju sendiri 5 langkah",
		Description: "player mu maju 5 langkah",
		EffectType:  "move_n_step",
		Target:      "self",
		Value:       5,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MAJU_2_LANGKAH_SELF = model.EffectConsumableItem{
		Name:        "maju sendiri 2 langkah",
		Description: "player mu maju 2 langkah",
		EffectType:  "move_n_step",
		Target:      "self",
		Value:       2,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MAJU_3_LANGKAH_SELF = model.EffectConsumableItem{
		Name:        "maju sendiri 3 langkah",
		Description: "player mu maju 3 langkah",
		EffectType:  "move_n_step",
		Target:      "self",
		Value:       3,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MAJU_1_LANGKAH_ALL = model.EffectConsumableItem{
		Name:        "dorong orang 1 langkah",
		Description: "player yang kamu pilih maju 1 langkah",
		EffectType:  "move_n_step",
		Target:      "all",
		Value:       1,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MAJU_2_LANGKAH_ALL = model.EffectConsumableItem{
		Name:        "dorong orang 2 langkah",
		Description: "player yang kamu pilih maju 2 langkah",
		EffectType:  "move_n_step",
		Target:      "all",
		Value:       2,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MAJU_3_LANGKAH_ALL = model.EffectConsumableItem{
		Name:        "dorong orang 3 langkah",
		Description: "player yang kamu pilih maju 3 langkah",
		EffectType:  "move_n_step",
		Target:      "all",
		Value:       3,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MAJU_5_LANGKAH_ALL = model.EffectConsumableItem{
		Name:        "dorong orang 5 langkah",
		Description: "player yang kamu pilih maju 5 langkah",
		EffectType:  "move_n_step",
		Target:      "all",
		Value:       5,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MUNDUR_10_LANGKAH_ALL = model.EffectConsumableItem{
		Name:        "tarik orang 10 langkah",
		Description: "player yang kamu pilih akan mundur 10 langkah",
		EffectType:  "move_n_step",
		Target:      "all",
		Value:       10,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}

	CONSUMABLE_MUNDUR_1_LANGKAH_ALL = model.EffectConsumableItem{
		Name:        "tarik orang 1 langkah",
		Description: "player yang kamu pilih akan mundur 1 langkah",
		EffectType:  "move_n_step",
		Target:      "all",
		Value:       1,
		UsePeriodMap: map[string]bool{
			"before_rolling_number": true,
			"before_moving":         true,
			"before_end_turn":       true,
		},
	}
)
