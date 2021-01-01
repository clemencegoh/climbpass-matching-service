package exceptions

import (
	"fmt"
)

// UserNotFoundException for handling Authentication errors
func UserNotFoundException() error {
	return fmt.Errorf("user does not exist or wrong credentials")
}

// UserAlreadyExists for handling Existing
func UserAlreadyExists() error {
	return fmt.Errorf("user already exists")
}
