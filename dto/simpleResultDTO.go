package dto

// SimpleResutlJSON is a struct for simple result in JSON format.x
type SimpleResutlJSON struct {
	Status    bool `json:"status"`
	ErrorCode int  `json:"error_code"`
}

// LoginResult is a struct for the result of login.
type LoginResult struct {
	Status    bool   `json:"status"`
	ErrorCode int    `json:"error_code"`
	Token     string `json:"token"`
}
