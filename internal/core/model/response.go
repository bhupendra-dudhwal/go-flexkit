package model

type Response struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Payload any    `json:"payload,omitempty"`
	Message string `json:"message"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Other   any    `json:"other"`
}
