package utils

import (
	"fmt"
	"strings"
)

func CheckMissingArgsForSelect(args []string) error {
	expectedStructure := []string{"SELECT", "<columns>", "FROM", "<table>"}

	if len(args) < len(expectedStructure) {
		for i, part := range args {
			if (i == 1 && part == "FROM") || (i == 2 && part != "FROM") {
				return fmt.Errorf("missing argument: %s", expectedStructure[i])
			}
		}
		return fmt.Errorf("missing argument: %s", expectedStructure[len(args)])
	}

	for i, part := range expectedStructure {
		if i >= len(args) || strings.TrimSpace(args[i]) == "" {
			return fmt.Errorf("missing argument: %s", part)
		}
		if (i == 0 && strings.ToUpper(args[i]) != "SELECT") || (i == 2 && strings.ToUpper(args[i]) != "FROM") {
			return fmt.Errorf("expected %s, got %s", part, args[i])
		}
	}

	return nil
}
