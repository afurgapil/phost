package utils

import (
	"testing"
)

func TestGetCommandType(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name            string
		args            args
		wantCommandType int
		wantErr         bool
	}{
		{
			name:            "Valid SELECT command",
			args:            args{command: "SELECT * FROM table"},
			wantCommandType: 0,
			wantErr:         false,
		},
		{
			name:            "Valid INSERT command",
			args:            args{command: "INSERT INTO table (col1, col2) VALUES (val1, val2)"},
			wantCommandType: 1,
			wantErr:         false,
		},
		{
			name:            "Valid DELETE command",
			args:            args{command: "DELETE FROM table WHERE id = 1"},
			wantCommandType: 2,
			wantErr:         false,
		},
		{
			name:            "Invalid command",
			args:            args{command: "UPDATE table SET col1 = val1"},
			wantCommandType: -1,
			wantErr:         true,
		},
		{
			name:            "Command with different case",
			args:            args{command: "select * from table"},
			wantCommandType: 0,
			wantErr:         false,
		},
		{
			name:            "Empty command",
			args:            args{command: ""},
			wantCommandType: -1,
			wantErr:         true,
		},
		{
			name:            "Command with trailing spaces",
			args:            args{command: "DELETE FROM table WHERE id = 1   "},
			wantCommandType: 2,
			wantErr:         false,
		},
		{
			name:            "Command with mixed case",
			args:            args{command: "DeLeTe FROM table WHERE id = 1"},
			wantCommandType: 2,
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCommandType, err := GetCommandType(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommandType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCommandType != tt.wantCommandType {
				t.Errorf("GetCommandType() = %v, want %v", gotCommandType, tt.wantCommandType)
			}
		})
	}
}
