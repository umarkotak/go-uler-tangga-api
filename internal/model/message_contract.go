package model

type (
	MessageContract struct {
		Action  string         `json:"action"`
		From    Identity       `json:"from"`
		To      Identity       `json:"to"`
		Payload MessagePayload `json:"payload"`
	}

	MessagePayload struct {
		ItemRandomID     string `json:"item_random_id"`
		ItemTargetUserID string `json:"item_target_user_id"`
	}
)
