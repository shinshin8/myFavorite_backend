package dto

// Profile saves user profile data.
type Profile struct {
	Status      int    `json:"status"`
	ErrorCode   int    `json:"error_code"`
	UserName    string `json:"user_name"`
	Birthday    string `json:"birthday"`
	MailAddress string `json:"mail_address"`
	Comment     string `json:"comment"`
}
