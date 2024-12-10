// Code generated by ent, DO NOT EDIT.

package room

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeDoors holds the string denoting the doors edge name in mutations.
	EdgeDoors = "doors"
	// EdgeDoorsIn holds the string denoting the doors_in edge name in mutations.
	EdgeDoorsIn = "doors_in"
	// EdgePlayers holds the string denoting the players edge name in mutations.
	EdgePlayers = "players"
	// Table holds the table name of the room in the database.
	Table = "rooms"
	// DoorsTable is the table that holds the doors relation/edge.
	DoorsTable = "doors"
	// DoorsInverseTable is the table name for the Door entity.
	// It exists in this package in order to avoid circular dependency with the "door" package.
	DoorsInverseTable = "doors"
	// DoorsColumn is the table column denoting the doors relation/edge.
	DoorsColumn = "door_from"
	// DoorsInTable is the table that holds the doors_in relation/edge.
	DoorsInTable = "doors"
	// DoorsInInverseTable is the table name for the Door entity.
	// It exists in this package in order to avoid circular dependency with the "door" package.
	DoorsInInverseTable = "doors"
	// DoorsInColumn is the table column denoting the doors_in relation/edge.
	DoorsInColumn = "door_to"
	// PlayersTable is the table that holds the players relation/edge.
	PlayersTable = "players"
	// PlayersInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayersInverseTable = "players"
	// PlayersColumn is the table column denoting the players relation/edge.
	PlayersColumn = "room_players"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Room queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByDoorsCount orders the results by doors count.
func ByDoorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDoorsStep(), opts...)
	}
}

// ByDoors orders the results by doors terms.
func ByDoors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDoorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDoorsInCount orders the results by doors_in count.
func ByDoorsInCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDoorsInStep(), opts...)
	}
}

// ByDoorsIn orders the results by doors_in terms.
func ByDoorsIn(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDoorsInStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlayersCount orders the results by players count.
func ByPlayersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlayersStep(), opts...)
	}
}

// ByPlayers orders the results by players terms.
func ByPlayers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDoorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DoorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, DoorsTable, DoorsColumn),
	)
}
func newDoorsInStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DoorsInInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, DoorsInTable, DoorsInColumn),
	)
}
func newPlayersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PlayersTable, PlayersColumn),
	)
}
