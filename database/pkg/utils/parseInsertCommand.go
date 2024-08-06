package utils

import "strings"

func ParseInsertCommand(command string) (args []string) {
	parts := strings.Fields(command)
	valuesPart := strings.Join(parts[4:], " ")
	if len(valuesPart) > 0 && valuesPart[0] == '(' && valuesPart[len(valuesPart)-1] == ')' {
		values := valuesPart[1 : len(valuesPart)-1]
		valueItems := strings.Split(values, ",")

		args = append(parts[:4], valueItems...)
		for i, v := range args[4:] {
			args[4+i] = strings.TrimSpace(v)
		}
		return args
	}
	return args

}
