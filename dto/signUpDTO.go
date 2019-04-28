package dto

// SignUpDTO struct
type SignUpDTO struct {
	Username     string
	EmailAddress string
	Password     string
}

// SignUpInfo create sign-up information DTO
// Username is in the first parameter.
// Email address is in the second parameter.
// Password is in the third parameter.
// It returns the instance of SignUpDTO.
func SignUpInfo(username string, emailAddress string, password string) *SignUpDTO {
	dto := new(SignUpDTO)
	dto.Username = username
	dto.EmailAddress = emailAddress
	dto.Password = password
	return dto
}
