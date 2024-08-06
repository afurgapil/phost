package parser

import (
	"errors"

	"github.com/afurgapil/phost/database/pkg/entities"
)

func ParseCommandInsert(command string) (entities.Command, error) {
	args, _, err := parseArgs(command)
	if err != nil {
		return entities.Command{}, err

	}

	if len(args) > 0 && args[len(args)-1] == "''" {
		return entities.Command{}, errors.New("invalid command: empty value detected")
	}

	return entities.Command{Type: entities.Insert, Args: args}, nil

}
