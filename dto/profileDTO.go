package dto

// ProfileResult saves user profile result.
type ProfileResult struct {
	Status    int `json:"status"`
	ErrorCode int `json:"error_code"`
	Profile   `json:"profile"`
}

// Profile saves user profile data.
type Profile struct {
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	Birthday    string `json:"birthday"`
	MailAddress string `json:"mail_address"`
	Comment     string `json:"comment"`
}
