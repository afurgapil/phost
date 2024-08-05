package parser

import (
	"reflect"
	"testing"
)

// TODO Test command could be extend
// TODO Check case sensetive
func TestParseCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		id         int
		name       string
		args       args
		want       Command
		wantErr    bool
		errMessage string
	}{
		{
			id:   1,
			name: "Valid SELECT command",
			args: args{command: "SELECT * FROM table"},
			want: Command{
				Type:        Select,
				Args:        []string{"SELECT", "*", "FROM", "table"},
				WhereClause: "",
			},
			wantErr: false,
		},
		{
			id:         2,
			name:       "Invalid SELECT command with missing table name",
			args:       args{command: "SELECT * FROM"},
			want:       Command{},
			wantErr:    true,
			errMessage: "missing argument: <table>",
		},
		{
			id:         3,
			name:       "Invalid SELECT command with missing columns",
			args:       args{command: "SELECT FROM table"},
			want:       Command{},
			wantErr:    true,
			errMessage: "missing argument: <columns>",
		},
		{
			id:         4,
			name:       "Invalid SELECT command with FROM",
			args:       args{command: "SELECT * table"},
			want:       Command{},
			wantErr:    true,
			errMessage: "missing argument: FROM",
		},
		{
			id:   5,
			name: "Valid SELECT command with WHERE",
			args: args{command: "SELECT * FROM table WHERE id=1"},
			want: Command{
				Type:        Select,
				Args:        []string{"SELECT", "*", "FROM", "table"},
				WhereClause: "id=1",
			},
			wantErr: false,
		},
		{
			id:         6,
			name:       "Invalid SELECT command with WHERE",
			args:       args{command: "SELECT * FROM table WHERE id="},
			want:       Command{},
			wantErr:    true,
			errMessage: "incomplete WHERE clause",
		},
		{
			id:      7,
			name:    "Valid INSERT command",
			args:    args{command: "INSERT INTO table VALUES ('value1')"},
			want:    Command{Type: Insert, Args: []string{"INSERT", "INTO", "table", "VALUES", "'value1'"}},
			wantErr: false,
		},
		{
			id:         8,
			name:       "Invalid INSERT command with null variable",
			args:       args{command: "INSERT INTO table VALUES ('')"},
			want:       Command{},
			wantErr:    true,
			errMessage: "invalid command: empty value detected",
		},
		{
			id:   9,
			name: "Valid DELETE command",
			args: args{command: "DELETE FROM table WHERE id=1"},
			want: Command{
				Type:        Delete,
				Args:        []string{"DELETE", "FROM", "table"},
				WhereClause: "id=1",
			},
			wantErr: false,
		},
		{
			id:         10,
			name:       "Invalid command",
			args:       args{command: "UPDATE table SET column=value"},
			want:       Command{},
			wantErr:    true,
			errMessage: "invalid command",
		},
		{
			id:         11,
			name:       "Empty command",
			args:       args{command: ""},
			want:       Command{},
			wantErr:    true,
			errMessage: "invalid command",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommand(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
			if err != nil && err.Error() != tt.errMessage {
				t.Errorf("ParseCommand() error message = %v, wantErrMessage %v", err.Error(), tt.errMessage)
			}
		})
	}
}

func Test_parseArgs(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		id              int
		name            string
		args            args
		wantArgs        []string
		wantWhereClause string
		wantErr         bool
		wantErrMessage  string
	}{
		{
			id:              1,
			name:            "Parse SELECT command",
			args:            args{command: "SELECT * FROM table WHERE id=1"},
			wantArgs:        []string{"SELECT", "*", "FROM", "table"},
			wantWhereClause: "id=1",
			wantErr:         false,
		},
		{
			id:              2,
			name:            "Parse INSERT command without WHERE",
			args:            args{command: "INSERT INTO table VALUES ('value1')"},
			wantArgs:        []string{"INSERT", "INTO", "table", "VALUES", "'value1'"},
			wantWhereClause: "",
			wantErr:         false,
		},
		{
			id:              3,
			name:            "Parse DELETE command with WHERE",
			args:            args{command: "DELETE FROM table WHERE id=1"},
			wantArgs:        []string{"DELETE", "FROM", "table"},
			wantWhereClause: "id=1",
			wantErr:         false,
		},
		{
			id:              4,
			name:            "Parse command with no WHERE clause",
			args:            args{command: "SELECT * FROM table"},
			wantArgs:        []string{"SELECT", "*", "FROM", "table"},
			wantWhereClause: "",
			wantErr:         false,
		},
		{
			id:              5,
			name:            "Parse empty command",
			args:            args{command: ""},
			wantArgs:        nil,
			wantWhereClause: "",
			wantErr:         true,
			wantErrMessage:  "missing argument",
		},
		{
			id:              6,
			name:            "Incomplete WHERE clause",
			args:            args{command: "SELECT * FROM table WHERE id="},
			wantArgs:        nil,
			wantWhereClause: "",
			wantErr:         true,
			wantErrMessage:  "incomplete WHERE clause",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArgs, gotWhereClause, err := parseArgs(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErrMessage != "" && err.Error() != tt.wantErrMessage {
				t.Errorf("parseArgs() error message = %v, wantErrMessage %v", err.Error(), tt.wantErrMessage)
			}
			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("parseArgs() gotArgs = %v, want %v", gotArgs, tt.wantArgs)
			}
			if gotWhereClause != tt.wantWhereClause {
				t.Errorf("parseArgs() gotWhereClause = %v, want %v", gotWhereClause, tt.wantWhereClause)
			}
		})
	}
}
