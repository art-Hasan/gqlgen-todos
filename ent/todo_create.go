// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/art-Hasan/gqlgen-todos/ent/todo"
	"github.com/art-Hasan/gqlgen-todos/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// TodoCreate is the builder for creating a Todo entity.
type TodoCreate struct {
	config
	mutation *TodoMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (tc *TodoCreate) SetCreatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (tc *TodoCreate) SetNillableCreatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the updated_at field.
func (tc *TodoCreate) SetUpdatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (tc *TodoCreate) SetNillableUpdatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetText sets the text field.
func (tc *TodoCreate) SetText(s string) *TodoCreate {
	tc.mutation.SetText(s)
	return tc
}

// SetDone sets the done field.
func (tc *TodoCreate) SetDone(b bool) *TodoCreate {
	tc.mutation.SetDone(b)
	return tc
}

// SetNillableDone sets the done field if the given value is not nil.
func (tc *TodoCreate) SetNillableDone(b *bool) *TodoCreate {
	if b != nil {
		tc.SetDone(*b)
	}
	return tc
}

// SetID sets the id field.
func (tc *TodoCreate) SetID(i int) *TodoCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetUserID sets the user edge to User by id.
func (tc *TodoCreate) SetUserID(id int) *TodoCreate {
	tc.mutation.SetUserID(id)
	return tc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (tc *TodoCreate) SetNillableUserID(id *int) *TodoCreate {
	if id != nil {
		tc = tc.SetUserID(*id)
	}
	return tc
}

// SetUser sets the user edge to User.
func (tc *TodoCreate) SetUser(u *User) *TodoCreate {
	return tc.SetUserID(u.ID)
}

// Save creates the Todo in the database.
func (tc *TodoCreate) Save(ctx context.Context) (*Todo, error) {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := todo.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := todo.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.Text(); !ok {
		return nil, errors.New("ent: missing required field \"text\"")
	}
	if v, ok := tc.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"text\": %v", err)
		}
	}
	if _, ok := tc.mutation.Done(); !ok {
		v := todo.DefaultDone
		tc.mutation.SetDone(v)
	}
	var (
		err  error
		node *Todo
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TodoCreate) SaveX(ctx context.Context) *Todo {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TodoCreate) sqlSave(ctx context.Context) (*Todo, error) {
	var (
		t     = &Todo{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: todo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todo.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		t.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCreatedAt,
		})
		t.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldUpdatedAt,
		})
		t.UpdatedAt = value
	}
	if value, ok := tc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldText,
		})
		t.Text = value
	}
	if value, ok := tc.mutation.Done(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: todo.FieldDone,
		})
		t.Done = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if t.ID == 0 {
		id := _spec.ID.Value.(int64)
		t.ID = int(id)
	}
	return t, nil
}
