package model

const (
	BROADCAST_ALL       = "all"
	BROADCAST_ROOM      = "room"
	BROADCAST_DIRECT_TO = "direct_to"
	BROADCAST_SELF      = "self"
)

type (
	ResponseContract struct {
		ResponseKind  string      `json:"response_kind"`
		BroadcastMode string      `json:"broadcast_mode"`
		From          Identity    `json:"from"`
		To            Identity    `json:"to"`
		Data          interface{} `json:"data"`
		ServerError   ServerError `json:"server_error"`
	}

	ServerError struct {
		Code          string `json:"code"`
		Message       string `json:"message"`
		CustomMessage string `json:"custom_message"`
	}
)
