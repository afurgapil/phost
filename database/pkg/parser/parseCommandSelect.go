package parser

import (
	"github.com/afurgapil/phost/database/pkg/entities"
	"github.com/afurgapil/phost/database/pkg/utils"
)

func ParseCommandSelect(command string) (entities.Command, error) {
	args, whereClause, err := parseArgs(command)
	if err != nil {
		return entities.Command{}, err
	}

	if err := utils.CheckMissingArgsForSelect(args); err != nil {
		return entities.Command{}, err
	}

	return entities.Command{Type: entities.Select, Args: args, WhereClause: whereClause}, nil
}
