package utils

import (
	"errors"
	"strings"
)

// ParseWhereIndex extracts the WHERE clause from a SQL command and returns
// the updated command without the WHERE clause. It also returns an error if
// the WHERE clause is incomplete or invalid.
func ParseWhereIndex(command string) (updatedCommand string, whereClause string, err error) {
	whereIndex := strings.Index(strings.ToUpper(command), "WHERE")
	if whereIndex != -1 {
		whereClause = strings.TrimSpace(command[whereIndex+len("WHERE"):])

		// Check if the WHERE clause is valid
		if len(whereClause) == 0 || !strings.Contains(whereClause, "=") {
			return "", "", errors.New("incomplete or invalid WHERE clause")
		}

		// Further check if the part after '=' is not empty
		whereParts := strings.Split(whereClause, "=")
		if len(whereParts) < 2 || len(strings.TrimSpace(whereParts[1])) == 0 {
			return "", "", errors.New("incomplete WHERE clause")
		}

		updatedCommand = strings.TrimSpace(command[:whereIndex])
		return updatedCommand, whereClause, nil
	}

	// No WHERE clause found
	return command, "", nil
}
