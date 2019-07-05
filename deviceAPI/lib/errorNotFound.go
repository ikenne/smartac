package lib

import "fmt"

type ErrorKeyNotFound struct {
	Key     string `json:"Id"`
	Message string `json:"message"`
}

func (e ErrorKeyNotFound) Error() string {
	return fmt.Sprintf("%s (Id=%s)", e.Message, e.Key)
}
