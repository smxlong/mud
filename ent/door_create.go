// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/smxlong/mud/ent/door"
	"github.com/smxlong/mud/ent/room"
)

// DoorCreate is the builder for creating a Door entity.
type DoorCreate struct {
	config
	mutation *DoorMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DoorCreate) SetName(s string) *DoorCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DoorCreate) SetDescription(s string) *DoorCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dc *DoorCreate) SetNillableDescription(s *string) *DoorCreate {
	if s != nil {
		dc.SetDescription(*s)
	}
	return dc
}

// SetDirection sets the "direction" field.
func (dc *DoorCreate) SetDirection(d door.Direction) *DoorCreate {
	dc.mutation.SetDirection(d)
	return dc
}

// SetID sets the "id" field.
func (dc *DoorCreate) SetID(s string) *DoorCreate {
	dc.mutation.SetID(s)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DoorCreate) SetNillableID(s *string) *DoorCreate {
	if s != nil {
		dc.SetID(*s)
	}
	return dc
}

// SetFromID sets the "from" edge to the Room entity by ID.
func (dc *DoorCreate) SetFromID(id string) *DoorCreate {
	dc.mutation.SetFromID(id)
	return dc
}

// SetNillableFromID sets the "from" edge to the Room entity by ID if the given value is not nil.
func (dc *DoorCreate) SetNillableFromID(id *string) *DoorCreate {
	if id != nil {
		dc = dc.SetFromID(*id)
	}
	return dc
}

// SetFrom sets the "from" edge to the Room entity.
func (dc *DoorCreate) SetFrom(r *Room) *DoorCreate {
	return dc.SetFromID(r.ID)
}

// SetToID sets the "to" edge to the Room entity by ID.
func (dc *DoorCreate) SetToID(id string) *DoorCreate {
	dc.mutation.SetToID(id)
	return dc
}

// SetNillableToID sets the "to" edge to the Room entity by ID if the given value is not nil.
func (dc *DoorCreate) SetNillableToID(id *string) *DoorCreate {
	if id != nil {
		dc = dc.SetToID(*id)
	}
	return dc
}

// SetTo sets the "to" edge to the Room entity.
func (dc *DoorCreate) SetTo(r *Room) *DoorCreate {
	return dc.SetToID(r.ID)
}

// Mutation returns the DoorMutation object of the builder.
func (dc *DoorCreate) Mutation() *DoorMutation {
	return dc.mutation
}

// Save creates the Door in the database.
func (dc *DoorCreate) Save(ctx context.Context) (*Door, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DoorCreate) SaveX(ctx context.Context) *Door {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DoorCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DoorCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DoorCreate) defaults() {
	if _, ok := dc.mutation.ID(); !ok {
		v := door.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DoorCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Door.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := door.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Door.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Direction(); !ok {
		return &ValidationError{Name: "direction", err: errors.New(`ent: missing required field "Door.direction"`)}
	}
	if v, ok := dc.mutation.Direction(); ok {
		if err := door.DirectionValidator(v); err != nil {
			return &ValidationError{Name: "direction", err: fmt.Errorf(`ent: validator failed for field "Door.direction": %w`, err)}
		}
	}
	return nil
}

func (dc *DoorCreate) sqlSave(ctx context.Context) (*Door, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Door.ID type: %T", _spec.ID.Value)
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DoorCreate) createSpec() (*Door, *sqlgraph.CreateSpec) {
	var (
		_node = &Door{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(door.Table, sqlgraph.NewFieldSpec(door.FieldID, field.TypeString))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(door.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(door.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := dc.mutation.Direction(); ok {
		_spec.SetField(door.FieldDirection, field.TypeEnum, value)
		_node.Direction = value
	}
	if nodes := dc.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   door.FromTable,
			Columns: []string{door.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.door_from = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   door.ToTable,
			Columns: []string{door.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.door_to = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DoorCreateBulk is the builder for creating many Door entities in bulk.
type DoorCreateBulk struct {
	config
	err      error
	builders []*DoorCreate
}

// Save creates the Door entities in the database.
func (dcb *DoorCreateBulk) Save(ctx context.Context) ([]*Door, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Door, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DoorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DoorCreateBulk) SaveX(ctx context.Context) []*Door {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DoorCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DoorCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
