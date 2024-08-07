package parser

import (
	"errors"

	"github.com/afurgapil/phost/database/pkg/entities"
)

func ParseCommandInsert(command string) (entities.Command, error) {
	args, _, err := parseArgs(command)

	if args[1] != "INTO" {
		return entities.Command{}, errors.New("invalid command: missing value INTO")
	}
	if args[2] == "VALUES" {
		return entities.Command{}, errors.New("invalid command: missing table name")
	}
	if args[3] != "VALUES" {
		return entities.Command{}, errors.New("invalid command: missing value VALUES")
	}
	if err != nil {
		return entities.Command{}, err

	}

	if len(args) > 0 && args[len(args)-1] == "''" {
		return entities.Command{}, errors.New("invalid command: empty value detected")
	}

	return entities.Command{Type: entities.Insert, Args: args}, nil

}
