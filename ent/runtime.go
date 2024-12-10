// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/smxlong/mud/ent/door"
	"github.com/smxlong/mud/ent/entity"
	"github.com/smxlong/mud/ent/room"
	"github.com/smxlong/mud/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	doorMixin := schema.Door{}.Mixin()
	doorMixinFields0 := doorMixin[0].Fields()
	_ = doorMixinFields0
	doorFields := schema.Door{}.Fields()
	_ = doorFields
	// doorDescName is the schema descriptor for name field.
	doorDescName := doorMixinFields0[1].Descriptor()
	// door.NameValidator is a validator for the "name" field. It is called by the builders before save.
	door.NameValidator = doorDescName.Validators[0].(func(string) error)
	// doorDescID is the schema descriptor for id field.
	doorDescID := doorMixinFields0[0].Descriptor()
	// door.DefaultID holds the default value on creation for the id field.
	door.DefaultID = doorDescID.Default.(func() string)
	entityFields := schema.Entity{}.Fields()
	_ = entityFields
	// entityDescName is the schema descriptor for name field.
	entityDescName := entityFields[1].Descriptor()
	// entity.NameValidator is a validator for the "name" field. It is called by the builders before save.
	entity.NameValidator = entityDescName.Validators[0].(func(string) error)
	// entityDescID is the schema descriptor for id field.
	entityDescID := entityFields[0].Descriptor()
	// entity.DefaultID holds the default value on creation for the id field.
	entity.DefaultID = entityDescID.Default.(func() string)
	roomMixin := schema.Room{}.Mixin()
	roomMixinFields0 := roomMixin[0].Fields()
	_ = roomMixinFields0
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescName is the schema descriptor for name field.
	roomDescName := roomMixinFields0[1].Descriptor()
	// room.NameValidator is a validator for the "name" field. It is called by the builders before save.
	room.NameValidator = roomDescName.Validators[0].(func(string) error)
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomMixinFields0[0].Descriptor()
	// room.DefaultID holds the default value on creation for the id field.
	room.DefaultID = roomDescID.Default.(func() string)
}
