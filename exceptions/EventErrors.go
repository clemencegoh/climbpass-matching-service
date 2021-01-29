package exceptions

import (
	"fmt"
)

func EventExistsException() error {
	return fmt.Errorf("Event already exists")
}
