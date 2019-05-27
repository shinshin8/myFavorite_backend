package dto

// SignUpResultJSON struct
type SignUpResultJSON struct {
	Status       int    `json:"status"`
	ErrorCode    int    `json:"error_code"`
	Username     string `json:"username"`
	EmailAddress string `json:"email_address"`
}

// SignUpResult returns the pointer of struct; signUpResultJSON.
// Status is in the first paramter.
// ErrorCode is in the second parameter.
// Username is in third parameter.
// EmailAddress is in the forth parameter.
// Password is in the fifth parameter.
// ConfirmPassword is in the sixth parameter.
func SignUpResult(status int, errorCode int, username string, emailAddress string) *SignUpResultJSON {
	res := new(SignUpResultJSON)
	res.Status = status
	res.ErrorCode = errorCode
	res.Username = username
	res.EmailAddress = emailAddress
	return res
}
