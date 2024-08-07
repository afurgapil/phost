package parser

import (
	"reflect"
	"testing"

	"github.com/afurgapil/phost/database/pkg/entities"
)

func TestParseCommandInsert(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Command
		wantErr bool
	}{
		{
			name: "Valid INSERT command",
			args: args{
				command: "INSERT INTO table VALUES ('value')",
			},
			want: entities.Command{
				Type: entities.Insert,
				Args: []string{"INSERT", "INTO", "table", "VALUES", "'value'"},
			},
			wantErr: false,
		},
		{
			name: "Invalid INSERT command with empty value",
			args: args{
				command: "INSERT INTO table VALUES ('')",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid INSERT command with missing VALUES keyword",
			args: args{
				command: "INSERT INTO table ('value')",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid INSERT command with missing table name",
			args: args{
				command: "INSERT INTO VALUES ('value')",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid INSERT command with missing INTO keyword",
			args: args{
				command: "INSERT table VALUES ('value')",
			},
			want:    entities.Command{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommandInsert(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommandInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
