package world_service

import (
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

func GetRoomData(messageContract model.MessageContract) (model.ResponseContract, error) {
	world := singleton.GetWorld()
	myIdentity := messageContract.From

	return model.ResponseContract{
		ResponseKind:  "room_data",
		BroadcastMode: model.BROADCAST_SELF,
		To:            model.Identity{},
		Data:          world.RoomMap[myIdentity.RoomID],
	}, nil
}
