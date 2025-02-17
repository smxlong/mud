// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/smxlong/mud/ent/room"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges        RoomEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// Doors holds the value of the doors edge.
	Doors []*Door `json:"doors,omitempty"`
	// DoorsIn holds the value of the doors_in edge.
	DoorsIn []*Door `json:"doors_in,omitempty"`
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DoorsOrErr returns the Doors value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) DoorsOrErr() ([]*Door, error) {
	if e.loadedTypes[0] {
		return e.Doors, nil
	}
	return nil, &NotLoadedError{edge: "doors"}
}

// DoorsInOrErr returns the DoorsIn value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) DoorsInOrErr() ([]*Door, error) {
	if e.loadedTypes[1] {
		return e.DoorsIn, nil
	}
	return nil, &NotLoadedError{edge: "doors_in"}
}

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[2] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldID, room.FieldName, room.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case room.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case room.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Room.
// This includes values selected through modifiers, order, etc.
func (r *Room) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryDoors queries the "doors" edge of the Room entity.
func (r *Room) QueryDoors() *DoorQuery {
	return NewRoomClient(r.config).QueryDoors(r)
}

// QueryDoorsIn queries the "doors_in" edge of the Room entity.
func (r *Room) QueryDoorsIn() *DoorQuery {
	return NewRoomClient(r.config).QueryDoorsIn(r)
}

// QueryPlayers queries the "players" edge of the Room entity.
func (r *Room) QueryPlayers() *PlayerQuery {
	return NewRoomClient(r.config).QueryPlayers(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return NewRoomClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room
