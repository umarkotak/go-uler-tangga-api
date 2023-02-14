package world_service

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

type UseItemResponse struct {
	ItemEffectType string       `json:"item_effect_type"`
	MoveResponse   MoveResponse `json:"move_response"`
	PlayerItems    []model.Item `json:"player_items"`
}

func UseItem(messageContract model.MessageContract) (model.ResponseContract, error) {
	world := singleton.GetWorld()
	myIdentity := messageContract.From
	room, _ := world.RoomMap[myIdentity.RoomID]
	player, _ := room.PlayerMap[myIdentity.ID]

	if room.ActivePlayer.Identity.ID != myIdentity.ID {
		return model.RESP_NOT_YOUR_TURN, nil
	}

	if messageContract.Payload.ItemRandomID == "" {
		return model.RESP_MISSING_ITEM_RANDOM_ID, nil
	}

	item := model.Item{}
	consumableItem := model.EffectConsumableItem{}
	var itemIdx int
	for idx, searchItem := range player.Items {
		if searchItem.RandomID == messageContract.Payload.ItemRandomID {
			item = searchItem
			consumableItem = item.EffectConsumableItem
			itemIdx = idx
			break
		}
	}
	if item.RandomID == "" {
		return model.RESP_ITEM_NOT_FOUND, nil
	}

	if player.NextState == model.STATE_ROLLING_NUMBER && consumableItem.UsePeriodMap["before_rolling_number"] {
	} else if player.NextState == model.STATE_MOVING && consumableItem.UsePeriodMap["before_moving"] {
	} else if player.NextState == model.STATE_ROLLING_NUMBER && consumableItem.UsePeriodMap["before_end_turn"] {
	} else {
		logrus.WithFields(logrus.Fields{
			"next_state":      player.NextState,
			"consumable_item": consumableItem,
		}).Error("err_invalid_state")
		return model.RESP_INVALID_STATE, nil
	}

	var targetPlayer model.Player
	var ok bool
	if consumableItem.Target == "self" {
		targetPlayer = player
	} else if consumableItem.Target == "all" {
		targetPlayer, ok = room.PlayerMap[messageContract.Payload.ItemTargetUserID]
		if !ok {
			return model.RESP_MISSING_TARGET_USER, nil
		}
	} else {
		return model.RESP_BAD_REQUEST, nil
	}

	useItemResponse := UseItemResponse{
		ItemEffectType: consumableItem.EffectType,
	}
	if consumableItem.EffectType == "move_n_step" {
		// Remove used item logic
		tmpPlayer := room.PlayerMap[player.Identity.ID]
		tmpPlayer.Items = removeItem(tmpPlayer.Items, itemIdx)
		room.PlayerMap[player.Identity.ID] = tmpPlayer

		movingCount := consumableItem.Value
		movingCount, _ = targetPlayer.CalculateCurrentPosition(room.MapConfig, movingCount)

		useItemResponse.MoveResponse = MoveResponse{
			Player: targetPlayer,
			Number: movingCount,
		}
		if targetPlayer.CheckIsWinning(room.MapConfig) {
			useItemResponse.MoveResponse.Winner = targetPlayer.Identity.ID
			targetPlayer.Winning = true
			room.Winners = append(room.Winners, useItemResponse.MoveResponse.Winner)
			targetPlayer.WinningPosition = int64(len(room.Winners))
		}
		useItemResponse.PlayerItems = room.PlayerMap[player.Identity.ID].Items

		room.PlayerMap[targetPlayer.Identity.ID] = targetPlayer

		if consumableItem.Target == "self" {
			room.WriteMoveLog(fmt.Sprintf("%v menggunakan %v", player.Identity.ID, consumableItem.Name))
		} else if consumableItem.Target == "all" {
			room.WriteMoveLog(fmt.Sprintf("%v menggunakan %v kepada %v", player.Identity.ID, consumableItem.Name, targetPlayer.Identity.ID))
		}
	}

	world.RoomMap[room.ID] = room
	playerMapToRoomIndex(room.ID)

	return model.ResponseContract{
		ResponseKind:  "player_used_item",
		BroadcastMode: model.BROADCAST_ROOM,
		To:            model.Identity{},
		Data:          useItemResponse,
	}, nil
}

func removeItem(slice []model.Item, idx int) []model.Item {
	return append(slice[:idx], slice[idx+1:]...)
}
