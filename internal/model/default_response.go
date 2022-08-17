package model

var (
	RESP_UNAUTHORIZED = ResponseContract{
		ResponseKind:  "err_unauthorized",
		BroadcastMode: BROADCAST_SELF,
	}

	RESP_INVALID_STATE = ResponseContract{
		ResponseKind:  "err_invalid_state",
		BroadcastMode: BROADCAST_SELF,
	}
)
