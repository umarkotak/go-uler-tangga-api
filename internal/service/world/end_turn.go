package world_service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

func EndTurn(messageContract model.MessageContract) (model.ResponseContract, error) {
	world := singleton.GetWorld()
	myIdentity := messageContract.From
	room, _ := world.RoomMap[myIdentity.RoomID]
	player, _ := room.PlayerMap[myIdentity.ID]

	if room.ActivePlayer.Identity.ID != myIdentity.ID {
		// logrus.Infof("%+v", room.ActivePlayer)
		return model.RESP_UNAUTHORIZED, nil
	}

	if player.NextState != model.STATE_END_TURN {
		// logrus.Infof("%+v", player)
		return model.RESP_INVALID_STATE, nil
	}

	player.CurrentState = model.STATE_END_TURN
	player.NextState = model.STATE_WAITING

	// TODO: Implement map effect for receiving reward
	movingCount := int64(0)
	isFoundItem := false

	playerFieldNumber := room.MapConfig.Numbering[room.MapConfig.Direction[player.IndexPosition]]
	fieldEffect, effectFound := room.MapConfig.FieldEffect[fmt.Sprintf("%v", playerFieldNumber)]
	if effectFound {
		if fieldEffect.BenefitType == "player_move" {
			movingCount = fieldEffect.EffectPlayerMove.MoveCount

		} else if fieldEffect.BenefitType == "consumable_item" {
			lenList := len(fieldEffect.EffectConsumableItems)
			if lenList > 0 {
				randomIndex := rand.Intn(lenList)
				effectConsumableItem := fieldEffect.EffectConsumableItems[randomIndex]

				player.Items = append(player.Items, model.Item{
					RandomID:             fmt.Sprint(time.Now().UnixNano()),
					EffectConsumableItem: effectConsumableItem,
				})
				isFoundItem = true
			}
		}
	}

	movingCount, winning := player.CalculateCurrentPosition(room.MapConfig, movingCount)
	moveResponse := MoveResponse{
		Player:      player,
		Number:      movingCount,
		IsFoundItem: isFoundItem,
	}
	if winning {
		moveResponse.Winner = player.Identity.ID
		player.Winning = true
		room.Winners = append(room.Winners, moveResponse.Winner)
		player.WinningPosition = int64(len(room.Winners))
	}
	room.PlayerMap[player.Identity.ID] = player

	if movingCount != 0 {
		direction := "maju"
		if movingCount < 0 {
			direction = "mundur"
		}
		room.WriteMoveLog(fmt.Sprintf("%v terkena efek %v %v langkah", myIdentity.ID, direction, movingCount))
	}

	nextRoomPlayerIndex := player.Identity.RoomPlayerIndex
	nextRoomPlayerIndex += 1
	if nextRoomPlayerIndex > int(room.PlayerCount) {
		nextRoomPlayerIndex = 1
	}
	for _, nextPlayer := range room.PlayerMap {
		if nextPlayer.Identity.RoomPlayerIndex == int(nextRoomPlayerIndex) {
			nextPlayer.CurrentState = model.STATE_PLAYING
			nextPlayer.NextState = model.STATE_ROLLING_NUMBER
			room.PlayerMap[nextPlayer.Identity.ID] = nextPlayer
			room.ActivePlayer = nextPlayer
			break
		}
	}

	moveResponse.NextPlayer = room.ActivePlayer

	world.RoomMap[room.ID] = room
	playerMapToRoomIndex(room.ID)

	room.WriteMoveLog(fmt.Sprintf("%v menyelesaikan giliran", myIdentity.ID))

	return model.ResponseContract{
		ResponseKind:  "player_end_turn",
		BroadcastMode: model.BROADCAST_ROOM,
		To:            model.Identity{},
		Data:          moveResponse,
	}, nil
}
