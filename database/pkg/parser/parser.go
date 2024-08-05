package parser

import (
	"errors"
	"strings"

	"github.com/afurgapil/phost/database/pkg/utils"
)

type CommandType int

const (
	Select CommandType = iota
	Insert
	Delete
)

type Command struct {
	Type        CommandType
	Args        []string
	WhereClause string
}

func ParseCommand(command string) (Command, error) {
	command = strings.TrimSpace(command)

	if strings.HasPrefix(strings.ToUpper(command), "SELECT") {
		args, whereClause, err := parseArgs(command)
		if err != nil {
			return Command{}, err
		}
		if err := utils.CheckMissingArgsForSelect(args); err != nil {
			return Command{}, err
		}
		return Command{Type: Select, Args: args, WhereClause: whereClause}, nil
	}

	if strings.HasPrefix(strings.ToUpper(command), "INSERT") {
		args, _, err := parseArgs(command)
		if err != nil {
			return Command{}, err
		}
		if len(args) > 0 {
			lastElement := args[len(args)-1]
			if lastElement == "''" {
				return Command{}, errors.New("invalid command: empty value detected")
			}
		}
		return Command{Type: Insert, Args: args}, nil
	}

	if strings.HasPrefix(strings.ToUpper(command), "DELETE") {
		args, whereClause, err := parseArgs(command)
		if err != nil {
			return Command{}, err
		}
		return Command{Type: Delete, Args: args, WhereClause: whereClause}, nil
	}

	return Command{}, errors.New("invalid command")
}

func parseArgs(command string) (args []string, whereClause string, err error) {
	parts := strings.Fields(command)
	if len(parts) < 3 {
		return nil, "", errors.New("missing argument")
	}

	whereIndex := strings.Index(strings.ToUpper(command), "WHERE")
	if whereIndex != -1 {
		whereClause = strings.TrimSpace(command[whereIndex+len("WHERE"):])
		if whereClause == "" || !strings.Contains(whereClause, "=") || len(strings.Split(whereClause, "=")[1]) == 0 {
			return nil, "", errors.New("incomplete WHERE clause")
		}
		command = command[:whereIndex]
	}

	parts = strings.Fields(command)

	if len(parts) > 4 && strings.ToUpper(parts[0]) == "INSERT" && strings.ToUpper(parts[3]) == "VALUES" {
		valuesPart := strings.Join(parts[4:], " ")
		if len(valuesPart) > 0 && valuesPart[0] == '(' && valuesPart[len(valuesPart)-1] == ')' {
			values := valuesPart[1 : len(valuesPart)-1]
			valueItems := strings.Split(values, ",")

			args = append(parts[:4], valueItems...)
			for i, v := range args[4:] {
				args[4+i] = strings.TrimSpace(v)
			}
		} else {
			args = parts
		}
	} else {
		args = parts
	}

	return args, whereClause, nil
}
