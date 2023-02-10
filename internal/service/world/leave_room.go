package world_service

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

func LeaveRoom(messageContract model.MessageContract) (model.ResponseContract, error) {
	logrus.Infof("Leaving in: %v", messageContract.From.ID)

	world := singleton.GetWorld()
	myIdentity := messageContract.From
	room, _ := world.RoomMap[myIdentity.RoomID]
	player, playerFound := room.PlayerMap[myIdentity.ID]

	if !playerFound {
		return model.ResponseContract{}, fmt.Errorf("Player not found")
	}

	player.IsOnline = false

	player.CurrentState = model.STATE_WAITING
	player.NextState = model.STATE_PLAYING

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

	room.PlayerMap[player.Identity.ID] = player
	room.WriteMoveLog(fmt.Sprintf("%v keluar", myIdentity.ID))
	world.RoomMap[room.ID] = room
	playerMapToRoomIndex(room.ID)

	return model.ResponseContract{
		ResponseKind:  "player_leave_room",
		BroadcastMode: model.BROADCAST_ROOM,
		To:            model.Identity{},
		Data:          world.RoomMap[room.ID],
	}, nil
}
