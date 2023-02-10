package model

var (
	RESP_UNAUTHORIZED = ResponseContract{
		ResponseKind:  "err_unauthorized",
		BroadcastMode: BROADCAST_SELF,
	}

	RESP_NOT_YOUR_TURN = ResponseContract{
		ResponseKind:  "err_not_your_turn",
		BroadcastMode: BROADCAST_SELF,
	}

	RESP_INVALID_STATE = ResponseContract{
		ResponseKind:  "err_invalid_state",
		BroadcastMode: BROADCAST_SELF,
	}

	RESP_MISSING_ITEM_RANDOM_ID = ResponseContract{
		ResponseKind:  "err_missing_item_random_id",
		BroadcastMode: BROADCAST_SELF,
	}

	RESP_ITEM_NOT_FOUND = ResponseContract{
		ResponseKind:  "err_item_not_found",
		BroadcastMode: BROADCAST_SELF,
	}

	RESP_ROOM_IS_FULL = ResponseContract{
		ResponseKind:  "err_room_is_full",
		BroadcastMode: BROADCAST_SELF,
	}

	RESP_BAD_REQUEST = ResponseContract{
		ResponseKind:  "err_bad_request",
		BroadcastMode: BROADCAST_SELF,
	}
)
