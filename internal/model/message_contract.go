package model

type (
	MessageContract struct {
		Action  string         `json:"action"`
		From    Identity       `json:"from"`
		To      Identity       `json:"to"`
		Payload MessagePayload `json:"payload"`
	}

	MessagePayload struct {
	}
)
