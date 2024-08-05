package utils

import "testing"

func TestCheckMissingArgsForSelect(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid SELECT command",
			args:    args{args: []string{"SELECT", "*", "FROM", "table"}},
			wantErr: false,
		},
		{
			name:    "Missing table name",
			args:    args{args: []string{"SELECT", "*", "FROM"}},
			wantErr: true,
			errMsg:  "missing argument: <table>",
		},
		{
			name:    "Missing columns",
			args:    args{args: []string{"SELECT", "FROM", "table"}},
			wantErr: true,
			errMsg:  "missing argument: <columns>",
		},
		{
			name:    "Missing FROM keyword",
			args:    args{args: []string{"SELECT", "*", "table"}},
			wantErr: true,
			errMsg:  "missing argument: FROM",
		},
		{
			name:    "Empty command",
			args:    args{args: []string{}},
			wantErr: true,
			errMsg:  "missing argument: SELECT",
		},
		{
			name:    "Extra arguments after table",
			args:    args{args: []string{"SELECT", "*", "FROM", "table", "extra"}},
			wantErr: false,
		},
		{
			name:    "Valid SELECT with whitespace arguments",
			args:    args{args: []string{"SELECT", "  *  ", "FROM", "  table  "}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckMissingArgsForSelect(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckMissingArgsForSelect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("CheckMissingArgsForSelect() error message = %v, wantErrMessage %v", err.Error(), tt.errMsg)
			}
		})
	}
}
