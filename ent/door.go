// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/smxlong/mud/ent/door"
	"github.com/smxlong/mud/ent/room"
)

// Door is the model entity for the Door schema.
type Door struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Direction holds the value of the "direction" field.
	Direction door.Direction `json:"direction,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DoorQuery when eager-loading is set.
	Edges        DoorEdges `json:"edges"`
	door_from    *string
	door_to      *string
	selectValues sql.SelectValues
}

// DoorEdges holds the relations/edges for other nodes in the graph.
type DoorEdges struct {
	// From holds the value of the from edge.
	From *Room `json:"from,omitempty"`
	// To holds the value of the to edge.
	To *Room `json:"to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FromOrErr returns the From value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DoorEdges) FromOrErr() (*Room, error) {
	if e.From != nil {
		return e.From, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: room.Label}
	}
	return nil, &NotLoadedError{edge: "from"}
}

// ToOrErr returns the To value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DoorEdges) ToOrErr() (*Room, error) {
	if e.To != nil {
		return e.To, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: room.Label}
	}
	return nil, &NotLoadedError{edge: "to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Door) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case door.FieldID, door.FieldName, door.FieldDescription, door.FieldDirection:
			values[i] = new(sql.NullString)
		case door.ForeignKeys[0]: // door_from
			values[i] = new(sql.NullString)
		case door.ForeignKeys[1]: // door_to
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Door fields.
func (d *Door) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case door.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = value.String
			}
		case door.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case door.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		case door.FieldDirection:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field direction", values[i])
			} else if value.Valid {
				d.Direction = door.Direction(value.String)
			}
		case door.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field door_from", values[i])
			} else if value.Valid {
				d.door_from = new(string)
				*d.door_from = value.String
			}
		case door.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field door_to", values[i])
			} else if value.Valid {
				d.door_to = new(string)
				*d.door_to = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Door.
// This includes values selected through modifiers, order, etc.
func (d *Door) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryFrom queries the "from" edge of the Door entity.
func (d *Door) QueryFrom() *RoomQuery {
	return NewDoorClient(d.config).QueryFrom(d)
}

// QueryTo queries the "to" edge of the Door entity.
func (d *Door) QueryTo() *RoomQuery {
	return NewDoorClient(d.config).QueryTo(d)
}

// Update returns a builder for updating this Door.
// Note that you need to call Door.Unwrap() before calling this method if this Door
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Door) Update() *DoorUpdateOne {
	return NewDoorClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Door entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Door) Unwrap() *Door {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Door is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Door) String() string {
	var builder strings.Builder
	builder.WriteString("Door(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteString(", ")
	builder.WriteString("direction=")
	builder.WriteString(fmt.Sprintf("%v", d.Direction))
	builder.WriteByte(')')
	return builder.String()
}

// Doors is a parsable slice of Door.
type Doors []*Door
