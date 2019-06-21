package dto

// SimpleResutlJSON is a struct for simple result in JSON format.x
type SimpleResutlJSON struct {
	Status    bool `json:"status"`
	ErrorCode int  `json:"error_code"`
}

type LoginResult struct {
	Status    bool   `json:"status"`
	ErrorCode int    `json:"error_code"`
	Token     string `json:"token"`
}
