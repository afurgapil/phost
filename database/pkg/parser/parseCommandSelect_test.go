package parser

import (
	"reflect"
	"testing"

	"github.com/afurgapil/phost/database/pkg/entities"
)

func TestParseCommandSelect(t *testing.T) {
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
			name: "Valid SELECT command with WHERE clause",
			args: args{
				command: "SELECT * FROM table WHERE id = 1",
			},
			want: entities.Command{
				Type:        entities.Select,
				Args:        []string{"SELECT", "*", "FROM", "table"},
				WhereClause: "id = 1",
			},
			wantErr: false,
		},
		{
			name: "Valid SELECT command without WHERE clause",
			args: args{
				command: "SELECT * FROM table",
			},
			want: entities.Command{
				Type: entities.Select,
				Args: []string{"SELECT", "*", "FROM", "table"},
			},
			wantErr: false,
		},
		{
			name: "Invalid SELECT command with missing FROM",
			args: args{
				command: "SELECT * table WHERE id = 1",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid SELECT command with incomplete WHERE clause",
			args: args{
				command: "SELECT * FROM table WHERE id",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid SELECT command with missing arguments",
			args: args{
				command: "SELECT",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid SELECT command with extra arguments",
			args: args{
				command: "SELECT * FROM table WHERE id = 1 EXTRA",
			},
			want:    entities.Command{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommandSelect(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandSelect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommandSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
