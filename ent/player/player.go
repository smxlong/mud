// Code generated by ent, DO NOT EDIT.

package player

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the player type in the database.
	Label = "player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// EdgePlayerRoles holds the string denoting the player_roles edge name in mutations.
	EdgePlayerRoles = "player_roles"
	// Table holds the table name of the player in the database.
	Table = "players"
	// RoomTable is the table that holds the room relation/edge.
	RoomTable = "players"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_players"
	// PlayerRolesTable is the table that holds the player_roles relation/edge. The primary key declared below.
	PlayerRolesTable = "player_player_roles"
	// PlayerRolesInverseTable is the table name for the PlayerRole entity.
	// It exists in this package in order to avoid circular dependency with the "playerrole" package.
	PlayerRolesInverseTable = "player_roles"
)

// Columns holds all SQL columns for player fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldPassword,
	FieldEmail,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "players"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"room_players",
}

var (
	// PlayerRolesPrimaryKey and PlayerRolesColumn2 are the table columns denoting the
	// primary key for the player_roles relation (M2M).
	PlayerRolesPrimaryKey = []string{"player_id", "player_role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Player queries.
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

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByRoomField orders the results by room field.
func ByRoomField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoomStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlayerRolesCount orders the results by player_roles count.
func ByPlayerRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlayerRolesStep(), opts...)
	}
}

// ByPlayerRoles orders the results by player_roles terms.
func ByPlayerRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayerRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRoomStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoomInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
	)
}
func newPlayerRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayerRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PlayerRolesTable, PlayerRolesPrimaryKey...),
	)
}
