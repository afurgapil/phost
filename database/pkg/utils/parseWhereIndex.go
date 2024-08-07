package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ParseWhereIndex(command string) (updatedCommand string, whereClause string, err error) {
	parts := strings.Fields(command)

	if parts[0] == "DELETE" && parts[1] == "FROM" && parts[2] == "WHERE" {
		return "", "", errors.New("missing argument: table name")

	}

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
		//we check numeric values for where command. if we update id values, should we update that
		if _, err := strconv.Atoi(strings.TrimSpace(whereParts[1])); err != nil {
			return "", "", errors.New("WHERE clause should have a numeric value after '='")
		}

		updatedCommand = strings.TrimSpace(command[:whereIndex])
		return updatedCommand, whereClause, nil
	}

	return command, "", nil
}
