package ghin

import (
	"fmt"
)

type UserNotLoggedInError struct {
	Msg string
}

func NewUserNotLoggedInError(msg string) error {
	return &UserNotLoggedInError{Msg: msg}
}

func (e UserNotLoggedInError) Error() string {
	return fmt.Sprintf("user is not logged in: %q", e.Msg)
}
