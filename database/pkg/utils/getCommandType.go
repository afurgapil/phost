package utils

import (
	"errors"
	"strings"
)

func GetCommandType(command string) (CommandType int, err error) {
	switch {
	case strings.HasPrefix(strings.ToUpper(command), "SELECT"):
		return 0, nil
	case strings.HasPrefix(strings.ToUpper(command), "INSERT"):
		return 1, nil
	case strings.HasPrefix(strings.ToUpper(command), "DELETE"):
		return 2, nil
	default:
		return -1, errors.New("invalid command")
	}
}
