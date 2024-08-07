package parser

import (
	"errors"
	"strings"

	"github.com/afurgapil/phost/database/pkg/entities"
	"github.com/afurgapil/phost/database/pkg/utils"
)

func ParseCommand(command string) (entities.Command, error) {
	command = strings.TrimSpace(command)

	cmdType, err := utils.GetCommandType(command)
	if err != nil {
		return entities.Command{}, err
	}

	switch cmdType {
	case 0:
		return ParseCommandSelect(command)
	case 1:
		return ParseCommandInsert(command)
	case 2:
		return ParseCommandDelete(command)
	default:
		return entities.Command{}, errors.New("invalid command")
	}

}

func parseArgs(command string) (args []string, whereClause string, err error) {
	parts := strings.Fields(command)
	if len(parts) < 3 {
		return nil, "", errors.New("missing argument")
	}
	updatedCommand, whereClause, err := utils.ParseWhereIndex(command)
	if err != nil {
		return nil, "", err
	}

	parts = strings.Fields(updatedCommand)

	if len(parts) > 4 && strings.ToUpper(parts[0]) == "INSERT" && strings.ToUpper(parts[3]) == "VALUES" {
		args = utils.ParseInsertCommand(updatedCommand)
	} else {
		args = parts
	}

	return args, whereClause, nil
}
