package world_service

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/go-uler-tangga-api/internal/global_setting"
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

func JoinRoom(messageContract model.MessageContract) (model.ResponseContract, error) {
	logrus.Infof("Coming in: %v", messageContract.From.ID)

	world := singleton.GetWorld()
	myIdentity := messageContract.From

	room, found := world.RoomMap[myIdentity.RoomID]

	if !found {
		room = model.Room{
			ID:           myIdentity.RoomID,
			MapConfig:    global_setting.MAP_1,
			PlayerMap:    map[string]model.Player{},
			PlayerCount:  0,
			ActivePlayer: model.Player{},
			MoveLogs:     []model.MoveLog{},
		}
		world.RoomMap[room.ID] = room

		logrus.Infof("Room creation success: %v", room.ID)
	}

	player, playerFound := room.PlayerMap[myIdentity.ID]

	if !playerFound {
		if room.PlayerCount >= room.MapConfig.MaxPlayer {
			err := fmt.Errorf("Room is full")
			return model.RESP_ROOM_IS_FULL, err
		}

		room.PlayerCount += 1
		myIdentity.RoomPlayerIndex = int(room.PlayerCount)
		myIdentity.RoomPlayerIndexString = fmt.Sprintf("%v", room.PlayerCount)
		player = model.Player{
			Identity:       myIdentity,
			Avatar:         model.Avatar{},
			AvatarPosition: model.AvatarPosition{},
			IndexPosition:  0,
			MapPosition:    []model.Position{},
			IsOnline:       true,
			Status:         model.STATUS_HEALTHY,
			CurrentState:   model.STATE_WAITING,
			NextState:      model.STATE_PLAYING,
			MoveAvailable:  0,
			Items:          []model.Item{},
		}
		player.Init(room.MapConfig, global_setting.DEFAULT_AVATAR_CONFIGS)
		room.WriteMoveLog(fmt.Sprintf("%v bergabung ke dalam permainan", myIdentity.ID))

	} else {
		player.IsOnline = true
		room.WriteMoveLog(fmt.Sprintf("%v masuk kembali", myIdentity.ID))
	}

	if room.PlayerCount == 1 {
		player.CurrentState = model.STATE_PLAYING
		player.NextState = model.STATE_ROLLING_NUMBER
		room.ActivePlayer = player
	}

	room.PlayerMap[myIdentity.ID] = player
	world.RoomMap[room.ID] = room

	playerMapToRoomIndex(room.ID)

	return model.ResponseContract{
		ResponseKind:  "player_join_room",
		BroadcastMode: model.BROADCAST_ROOM,
		To:            model.Identity{},
		Data:          world.RoomMap[room.ID],
	}, nil
}

func playerMapToRoomIndex(roomID string) {
	world := singleton.GetWorld()
	room := world.RoomMap[roomID]

	room.PlayerRoomIndexMap = map[string]model.Player{}
	for _, player := range room.PlayerMap {
		room.PlayerRoomIndexMap[player.Identity.RoomPlayerIndexString] = player
	}

	world.RoomMap[roomID] = room
}
