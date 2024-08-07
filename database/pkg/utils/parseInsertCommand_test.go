package utils

import (
	"reflect"
	"testing"
)

func TestParseInsertCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name     string
		args     args
		wantArgs []string
	}{
		{
			name:     "Valid INSERT command with single value",
			args:     args{command: "INSERT INTO table VALUES ('value')"},
			wantArgs: []string{"INSERT", "INTO", "table", "VALUES", "'value'"},
		},
		{
			name:     "Valid INSERT command with value having spaces",
			args:     args{command: "INSERT INTO table VALUES ('value with spaces')"},
			wantArgs: []string{"INSERT", "INTO", "table", "VALUES", "'value with spaces'"},
		},
		{
			name:     "Invalid INSERT command with missing value",
			args:     args{command: "INSERT INTO table VALUES ('')"},
			wantArgs: []string{"INSERT", "INTO", "table", "VALUES", "''"},
		},
		{
			name:     "Invalid INSERT command with extra spaces",
			args:     args{command: "INSERT INTO   table   VALUES   ('value')"},
			wantArgs: []string{"INSERT", "INTO", "table", "VALUES", "'value'"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArgs := ParseInsertCommand(tt.args.command); !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("ParseInsertCommand() = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}
