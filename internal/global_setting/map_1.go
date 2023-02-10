package global_setting

import (
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
)

var (
	MAP_1 = model.MapConfig{
		Title:     "Default",
		MinNumber: 1,
		MaxNumber: 7,
		Size:      100,
		MaxPlayer: 5,
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
			"2": {
				FieldNumber:       2,
				FieldNumberString: "2",
				BenefitType:       "boost_item",
				EffectBoostItem:   model.EffectBoostItem{},
			},
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
			"6": {
				FieldNumber:       6,
				FieldNumberString: "6",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      8,
					FromCoordinate: 6,
					ToCoordinate:   14,
				},
			},
			"8": {
				FieldNumber:       8,
				FieldNumberString: "8",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      4,
					FromCoordinate: 8,
					ToCoordinate:   12,
				},
			},
			"11": {
				FieldNumber:       11,
				FieldNumberString: "11",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      17,
					FromCoordinate: 11,
					ToCoordinate:   28,
				},
			},
			"15": {
				FieldNumber:       15,
				FieldNumberString: "15",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      19,
					FromCoordinate: 15,
					ToCoordinate:   34,
				},
			},
			"18": {
				FieldNumber:       18,
				FieldNumberString: "18",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "down",
					MoveCount:      -17,
					FromCoordinate: 18,
					ToCoordinate:   1,
				},
			},
			"22": {
				FieldNumber:       22,
				FieldNumberString: "22",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      15,
					FromCoordinate: 22,
					ToCoordinate:   37,
				},
			},
			"26": {
				FieldNumber:       26,
				FieldNumberString: "26",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "down",
					MoveCount:      -16,
					FromCoordinate: 26,
					ToCoordinate:   10,
				},
			},
			"38": {
				FieldNumber:       38,
				FieldNumberString: "38",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      21,
					FromCoordinate: 38,
					ToCoordinate:   59,
				},
			},
			"39": {
				FieldNumber:       39,
				FieldNumberString: "39",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "down",
					MoveCount:      -34,
					FromCoordinate: 39,
					ToCoordinate:   5,
				},
			},
			"49": {
				FieldNumber:       49,
				FieldNumberString: "49",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "up",
					MoveCount:      18,
					FromCoordinate: 49,
					ToCoordinate:   67,
				},
			},
			"51": {
				FieldNumber:       51,
				FieldNumberString: "51",
				BenefitType:       "player_move",
				EffectPlayerMove: model.EffectPlayerMove{
					Direction:      "down",
					MoveCount:      -44,
					FromCoordinate: 51,
					ToCoordinate:   7,
				},
			},
		},
	}
)
