/*
	validation.go is handling validations.
*/

package utils

import (
	"regexp"
)

// IsName checks if name in parameter can be usable.
// It checks if name length is not 0 and 0 to 20.
// This returns true or false.
func IsName(name string) bool {
	if len(name) == 0 || len(name) <= 0 || len(name) >= 20 {
		return false
	}
	return true
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
	if len(password) == 0 || len(password) < 8 || len(password) > 20 {
		return false
	}
	return true
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
	return (len(title) == 0 || len(title) <= 0 || len(title) >= 51)
}

// IsContent checks if a value is valied content.
// It checks
// - a value is not empty.
// - a value's length is 1 to 500.
// It returns a result in boolean.
func IsContent(content string) bool {
	return (len(content) == 0 || len(content) <= 0 || len(content) >= 501)
}
