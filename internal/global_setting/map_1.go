package global_setting

import (
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
)

var (
	MAP_1 = model.MapConfig{
		MinNumber: 2,
		MaxNumber: 3,
		Size:      100,
		Numbering: []int64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			20, 19, 18, 17, 16, 15, 14, 13, 12, 11,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			40, 39, 38, 37, 36, 35, 34, 33, 32, 31,
			41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
			60, 59, 58, 57, 56, 55, 54, 53, 52, 51,
			61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
			80, 79, 78, 77, 76, 75, 74, 73, 72, 71,
			81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
			100, 99, 98, 97, 96, 95, 94, 93, 92, 91,
		},
		Direction: []int64{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			19, 18, 17, 16, 15, 14, 13, 12, 11, 10,
			20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
			39, 38, 37, 36, 35, 34, 33, 32, 31, 30,
			40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
			59, 58, 57, 56, 55, 54, 53, 52, 51, 50,
			60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
			79, 78, 77, 76, 75, 74, 73, 72, 71, 70,
			80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
			99, 98, 97, 96, 95, 94, 93, 92, 91, 90,
		},
		FieldEffect: map[string]model.FieldEffect{
			"3": {
				FieldNumber:       3,
				FieldNumberString: "3",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      17,
					FromCoordinate: 3,
					ToCoordinate:   20,
				},
			},
		},
	}
)
