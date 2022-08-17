package model

const (
	BROADCAST_ALL  = "all"
	BROADCAST_ROOM = "room"
	BROADCAST_ONE  = "one"
	BROADCAST_SELF = "self"
)

type (
	ResponseContract struct {
		ResponseKind  string      `json:"response_kind"`
		BroadcastMode string      `json:"broadcast_mode"`
		To            Identity    `json:"to"`
		Data          interface{} `json:"data"`
	}
)
