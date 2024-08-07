package parser

import (
	"github.com/afurgapil/phost/database/pkg/entities"
)

func ParseCommandDelete(command string) (entities.Command, error) {
	args, whereClause, err := parseArgs(command)
	if err != nil {
		return entities.Command{}, err
	}

	return entities.Command{Type: entities.Delete, Args: args, WhereClause: whereClause}, nil

}
