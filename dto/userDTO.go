package dto

// LoginBody is a struct for request body of login.
type LoginBody struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// SingUpBody is a sturct for request body of sign up.
type SingUpBody struct {
	UserName        string `json:"username"`
	EmailAddress    string `json:"mail_address"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// SignUpResultJSON struct
type SignUpResultJSON struct {
	Status       bool   `json:"status"`
	ErrorCode    int    `json:"error_code"`
	Username     string `json:"username"`
	EmailAddress string `json:"email_address"`
}

// EditProfileBody is a struct for request body of edit profile.
type EditProfileBody struct {
	UserName    string `json:"username"`
	Birthday    string `json:"birthday"`
	MailAddress string `json:"mail_address"`
	Comment     string `json:"comment"`
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

// DeleteImage is a strict for image urls' for deleting.
type DeleteImage struct {
	ImageURL []string `json:"delete_url"`
}
