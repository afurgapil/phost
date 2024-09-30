package utils

import (
	"errors"
	"strings"
)

func ParseWhereIndex(command string) (updatedCommand string, whereClause string, err error) {
	whereIndex := strings.Index(strings.ToUpper(command), "WHERE")
	if whereIndex != -1 {
		whereClause = strings.TrimSpace(command[whereIndex+len("WHERE"):])

		if len(whereClause) == 0 || !strings.Contains(whereClause, "=") {
			return "", "", errors.New("incomplete or invalid WHERE clause")
		}

		whereParts := strings.Split(whereClause, "=")
		if len(whereParts) < 2 || len(strings.TrimSpace(whereParts[1])) == 0 {
			return "", "", errors.New("incomplete WHERE clause")
		}

		updatedCommand = strings.TrimSpace(command[:whereIndex])
		return updatedCommand, whereClause, nil
	}

	return command, "", nil
}
