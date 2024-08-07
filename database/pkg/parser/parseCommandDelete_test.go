package parser

import (
	"reflect"
	"testing"

	"github.com/afurgapil/phost/database/pkg/entities"
)

func TestParseCommandDelete(t *testing.T) {
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
			name: "Valid DELETE command",
			args: args{
				command: "DELETE FROM table",
			},
			want: entities.Command{
				Type: entities.Delete,
				Args: []string{"DELETE", "FROM", "table"},
			},
			wantErr: false,
		},
		{
			name: "Valid DELETE command with WHERE",
			args: args{
				command: "DELETE FROM table WHERE id = 1",
			},
			want: entities.Command{
				Type:        entities.Delete,
				Args:        []string{"DELETE", "FROM", "table"},
				WhereClause: "id = 1",
			},
			wantErr: false,
		},
		{
			name: "Invalid DELETE command with missing table name",
			args: args{
				command: "DELETE FROM",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid DELETE command with missing table name with WHERE",
			args: args{
				command: "DELETE FROM WHERE id = 1",
			},
			want:    entities.Command{},
			wantErr: true,
		},
		{
			name: "Invalid DELETE command with incomplete WHERE clause",
			args: args{
				command: "DELETE FROM table WHERE id = ",
			},
			want:    entities.Command{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommandDelete(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommandDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
