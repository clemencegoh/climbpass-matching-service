package exceptions

import (
	"fmt"
)

func GymExistsException(name string) error {
	return fmt.Errorf("Gym with name %s already exists", name)
}
