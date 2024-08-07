package entities

type CommandType int

const (
	Select CommandType = iota
	Insert
	Delete
)

type Command struct {
	Type        CommandType
	Args        []string
	WhereClause string
}
