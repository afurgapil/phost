package parser

import (
	"errors"
	"strings"
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

	if strings.HasPrefix(command, "SELECT") {
		args, whereClause := parseArgs(command)
		return Command{Type: Select, Args: args, WhereClause: whereClause}, nil
	}
	if strings.HasPrefix(command, "INSERT") {
		args, _ := parseArgs(command)
		return Command{Type: Insert, Args: args}, nil
	}
	if strings.HasPrefix(command, "DELETE") {
		args, whereClause := parseArgs(command)
		return Command{Type: Delete, Args: args, WhereClause: whereClause}, nil
	}
	return Command{}, errors.New("invalid command")
}

func parseArgs(command string) (args []string, whereClause string) {
	parts := strings.Fields(command)
	if len(parts) < 3 {
		return nil, ""
	}

	whereIndex := strings.Index(command, "WHERE")
	if whereIndex != -1 {
		whereClause = strings.TrimSpace(command[whereIndex+len("WHERE"):])
		command = command[:whereIndex]
	}

	println("Command:", command, "where:", whereClause)
	return strings.Fields(command), whereClause
}
