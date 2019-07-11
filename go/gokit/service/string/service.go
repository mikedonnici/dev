package main

import (
	"errors"
	"strings"
)

// ErrEmpty is an error message for an empty string
var ErrEmpty = errors.New("Empty string")

// StringService does stuff with strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// Implementation
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if len(s) == 0 {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
