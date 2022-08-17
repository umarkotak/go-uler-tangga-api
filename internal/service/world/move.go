package world_service

import (
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

type MoveResponse struct {
	Player     model.Player `json:"player"`
	NextPlayer model.Player `json:"next_player"`
	Number     int64        `json:"number"`
}

func Move(messageContract model.MessageContract) (model.ResponseContract, error) {
	world := singleton.GetWorld()
	myIdentity := messageContract.From
	room, _ := world.RoomMap[myIdentity.RoomID]
	player, _ := room.PlayerMap[myIdentity.ID]

	if room.ActivePlayer.Identity.ID != myIdentity.ID {
		// logrus.Infof("%+v", room.ActivePlayer)
		return model.RESP_UNAUTHORIZED, nil
	}

	if player.NextState != model.STATE_MOVING {
		// logrus.Infof("%+v", player)
		return model.RESP_INVALID_STATE, nil
	}

	movingCount := player.MoveAvailable

	player.CurrentState = model.STATE_MOVING
	player.NextState = model.STATE_WAITING
	player.MoveAvailable = 0
	movingCount = player.CalculateCurrentPosition(room.MapConfig, movingCount)
	moveResponse := MoveResponse{
		Player: player,
		Number: movingCount,
	}
	room.PlayerMap[player.Identity.ID] = player

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

	return model.ResponseContract{
		ResponseKind:  "player_move",
		BroadcastMode: model.BROADCAST_ROOM,
		To:            model.Identity{},
		Data:          moveResponse,
	}, nil
}
