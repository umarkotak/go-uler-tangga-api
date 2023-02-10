package world_service

import (
	"fmt"
	"math/rand"

	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

type RollNumberResponse struct {
	Player model.Player `json:"player"`
	Number int64        `json:"number"`
}

func RollNumber(messageContract model.MessageContract) (model.ResponseContract, error) {
	world := singleton.GetWorld()
	myIdentity := messageContract.From
	room, _ := world.RoomMap[myIdentity.RoomID]
	player, _ := room.PlayerMap[myIdentity.ID]

	if room.ActivePlayer.Identity.ID != myIdentity.ID {
		// logrus.Infof("%+v", room.ActivePlayer)
		return model.RESP_UNAUTHORIZED, nil
	}

	if player.NextState != model.STATE_ROLLING_NUMBER {
		// logrus.Infof("%+v", player)
		return model.RESP_INVALID_STATE, nil
	}

	number := int64(rand.Intn(int(room.MapConfig.MaxNumber-room.MapConfig.MinNumber)) + int(room.MapConfig.MinNumber))

	player.CurrentState = model.STATE_ROLLING_NUMBER
	player.NextState = model.STATE_MOVING
	player.MoveAvailable = number
	rollNumberResponse := RollNumberResponse{
		Player: player,
		Number: player.MoveAvailable,
	}

	room.PlayerMap[player.Identity.ID] = player
	world.RoomMap[room.ID] = room
	playerMapToRoomIndex(room.ID)

	room.WriteMoveLog(fmt.Sprintf("%v mendapatkan angka %v", myIdentity.ID, rollNumberResponse.Player.MoveAvailable))

	return model.ResponseContract{
		ResponseKind:  "player_roll_number",
		BroadcastMode: model.BROADCAST_ROOM,
		To:            model.Identity{},
		Data:          rollNumberResponse,
	}, nil
}
