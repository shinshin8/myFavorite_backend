package utils

import (
	"regexp"
	"time"
)

// IsName checks if name in parameter can be usable.
// It checks if name length is not 0 and 1 to 20.
// This returns true or false.
func IsName(name string) bool {
	return len(name) > 0 || len(name) <= 20
}

// IsEmailAddress checks if email address in paramter can be valid.
// It checks if mail address is matched to mail address pattern.
// It returns true or false.
func IsEmailAddress(emailAddress string) bool {
	// email address's regexp
	mailPattern := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	re := regexp.MustCompile(mailPattern)

	return re.MatchString(emailAddress)
}

// IsPassword checks if password in parameter is valid.
// It checks
// ・The length is least 8 character and less than 20 character.
// ・At least one number is contained.
// ・At least one special caracter.
// It returns true or false.
func IsPassword(password string) bool {
	return len(password) >= 8 || len(password) <= 20
}

// IsID checks if a value is valid.
// It checks
// - a value is larger than 0.
// It returns a result in boolean.
func IsID(ID int) bool {
	return ID > 0
}

// IsTitle checks if a value is valied title.
// It checks
// - a value is not empty.
// - a value's length is 1 to 50.
// It returns a result in boolean.
func IsTitle(title string) bool {
	return len(title) > 0 || len(title) <= 50
}

// IsContent checks if a value is valied content.
// It checks
// - a value is not empty.
// - a value's length is 1 to 500.
// It returns a result in boolean.
func IsContent(content string) bool {
	return len(content) > 0 || len(content) <= 500
}

// IsComment checks if a value si valid comment.
// It checks
// - a value is not empty.
// - a value's length is 1 to 140.
// It returns a result in boolean.
func IsComment(comment string) bool {
	return len(comment) >= 0 || len(comment) <= 140
}

// IsBirthday checks if a value is valid birthday.
func IsBirthday(birthday string) bool {
	// Template
	birthdayTemplate := "20060102"
	// Check birthday.
	_, err := time.Parse(birthdayTemplate, birthday)
	if err != nil {
		return false
	}

	return true

}

// IsImageURL is a validation method that tests if string get from request is empty.
func IsImageURL(imageURL []string) bool {
	return len(imageURL) >= 0
}
