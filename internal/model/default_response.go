package model

var (
	RESP_UNAUTHORIZED = ResponseContract{
		ResponseKind:  "process_error",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_unauthorized",
			Message: "tidak ter authorisasi",
		},
	}

	RESP_NOT_YOUR_TURN = ResponseContract{
		ResponseKind:  "process_error",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_not_your_turn",
			Message: "bukan giliran mu",
		},
	}

	RESP_INVALID_STATE = ResponseContract{
		ResponseKind:  "process_error",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_invalid_state",
			Message: "state salah",
		},
	}

	RESP_MISSING_ITEM_RANDOM_ID = ResponseContract{
		ResponseKind: "	",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_missing_item_random_id",
			Message: "item belum dipilih",
		},
	}

	RESP_ITEM_NOT_FOUND = ResponseContract{
		ResponseKind:  "process_error",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_item_not_found",
			Message: "item tidak dapat ditemukan",
		},
	}

	RESP_ROOM_IS_FULL = ResponseContract{
		ResponseKind:  "process_error",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_room_is_full",
			Message: "room penuh",
		},
	}

	RESP_BAD_REQUEST = ResponseContract{
		ResponseKind:  "process_error",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_bad_request",
			Message: "bad request",
		},
	}

	RESP_MISSING_TARGET_USER = ResponseContract{
		ResponseKind:  "process_error",
		BroadcastMode: BROADCAST_SELF,
		ServerError: ServerError{
			Code:    "err_missing_target_user",
			Message: "target user belum dipilih",
		},
	}
)
