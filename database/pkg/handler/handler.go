package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/afurgapil/phost/database/internal/database"
	"github.com/afurgapil/phost/database/pkg/entities"
	"github.com/afurgapil/phost/database/pkg/parser"
)

var db *database.Database

func SetDatabase(database *database.Database) {
	db = database
}

func HandleExecute(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("command")
	if query == "" {
		http.Error(w, "Command is required", http.StatusBadRequest)
		return
	}

	command, err := parser.ParseCommand(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result interface{}
	switch entities.CommandType(command.Type) {
	case entities.Select:
		result, err = executeSelect(command.Args, command.WhereClause)
	case entities.Insert:
		id, recordedValue, err := executeInsert(command.Args)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = map[string]interface{}{"id": id, "value": recordedValue}
	case entities.Delete:
		err = executeDelete(command.Args, command.WhereClause)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result != nil {
		w.Header().Set("Content-Type", "application/json") // JSON içeriği belirle
		json.NewEncoder(w).Encode(result)
	}
}

func executeSelect(args []string, whereClause string) ([]database.Record, error) {
	if len(args) < 3 {
		return nil, errors.New("missing argument in SELECT command")
	}
	var result []database.Record
	if args[1] == "*" && args[2] == "FROM" && args[3] == "records" {
		if whereClause != "" {
			for _, record := range db.Records {
				if whereClause == "id="+strconv.Itoa(record.ID) {
					result = append(result, record)
				}
			}
		} else {
			result = db.Records
		}
	} else {
		return nil, errors.New("invalid SELECT command")
	}

	if len(result) == 0 {
		return nil, errors.New("no record found")
	}

	return result, nil
}

func executeInsert(args []string) (id int, recordedValue string, error error) {
	if len(args) < 5 || strings.ToUpper(args[3]) != "VALUES" {
		return 0, "", errors.New("INSERT command missing argument or keyword 'VALUES' not found")
	}

	valuesIndex := 4
	value := strings.Join(args[valuesIndex:], " ")

	if len(value) > 1 && value[0] == '\'' && value[len(value)-1] == '\'' {
		value = value[1 : len(value)-1]

	}

	if strings.TrimSpace(value) == "" {
		return 0, "", errors.New("no null value can be inserted in the INSERT command")
	}

	id, recordedValue = db.AddRecord(value)
	return id, recordedValue, nil
}

func executeDelete(args []string, whereClause string) error {
	if len(args) < 3 {
		return errors.New("DELETE command missing argument")
	}

	if whereClause == "" {
		db.ClearRecords()
		return nil
	}

	idStr := strings.TrimPrefix(whereClause, "id=")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("invalid ID format")
	}

	if !db.DeleteRecord(id) {
		return errors.New("could not delete record")
	}

	return nil
}
