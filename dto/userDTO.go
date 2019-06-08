package dto

// LoginBody is a struct for request body of login.
type LoginBody struct {
	UserName string
	Password string
}

// SingUpBody is a sturct for request body of sign up.
type SingUpBody struct {
	UserName        string
	Email           string
	Password        string
	ConfirmPassword string
}

// EditProfileBody is a struct for request body of edit profile.
type EditProfileBody struct {
	UserName    string
	Birthday    string
	MailAddress string
	Comment     string
}

// ProfileResult saves user profile result.
type ProfileResult struct {
	Status    bool `json:"status"`
	ErrorCode int  `json:"error_code"`
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
