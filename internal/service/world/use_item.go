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
	for _, searchItem := range player.Items {
		if searchItem.RandomID == messageContract.Payload.ItemRandomID {
			item = searchItem
			consumableItem = item.EffectConsumableItem
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
	if consumableItem.Target == "self" {
		targetPlayer = player
	} else {
		return model.RESP_BAD_REQUEST, nil
	}

	useItemResponse := UseItemResponse{
		ItemEffectType: consumableItem.EffectType,
	}
	if consumableItem.EffectType == "move_n_step" {
		movingCount := consumableItem.Value
		movingCount = targetPlayer.CalculateCurrentPosition(room.MapConfig, movingCount)

		useItemResponse.MoveResponse = MoveResponse{
			Player: targetPlayer,
			Number: movingCount,
		}
		room.PlayerMap[targetPlayer.Identity.ID] = targetPlayer

		room.WriteMoveLog(fmt.Sprintf("%v jalan %v langkah", targetPlayer.Identity.ID, consumableItem.Value))
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
