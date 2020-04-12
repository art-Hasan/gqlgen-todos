// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/art-Hasan/gqlgen-todos/ent/todo"
	"github.com/art-Hasan/gqlgen-todos/ent/user"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTodo = "Todo"
	TypeUser = "User"
)

// TodoMutation represents an operation that mutate the Todos
// nodes in the graph.
type TodoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	created_at    *time.Time
	updated_at    *time.Time
	text          *string
	done          *bool
	clearedFields map[string]struct{}
}

var _ ent.Mutation = (*TodoMutation)(nil)

// newTodoMutation creates new mutation for $n.Name.
func newTodoMutation(c config, op Op) *TodoMutation {
	return &TodoMutation{
		config:        c,
		op:            op,
		typ:           TypeTodo,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TodoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TodoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Todo creation.
func (m *TodoMutation) SetID(id int) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TodoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreatedAt sets the created_at field.
func (m *TodoMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *TodoMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt reset all changes of the created_at field.
func (m *TodoMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *TodoMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *TodoMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdatedAt reset all changes of the updated_at field.
func (m *TodoMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetText sets the text field.
func (m *TodoMutation) SetText(s string) {
	m.text = &s
}

// Text returns the text value in the mutation.
func (m *TodoMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// ResetText reset all changes of the text field.
func (m *TodoMutation) ResetText() {
	m.text = nil
}

// SetDone sets the done field.
func (m *TodoMutation) SetDone(b bool) {
	m.done = &b
}

// Done returns the done value in the mutation.
func (m *TodoMutation) Done() (r bool, exists bool) {
	v := m.done
	if v == nil {
		return
	}
	return *v, true
}

// ResetDone reset all changes of the done field.
func (m *TodoMutation) ResetDone() {
	m.done = nil
}

// Op returns the operation name.
func (m *TodoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Todo).
func (m *TodoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *TodoMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, todo.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, todo.FieldUpdatedAt)
	}
	if m.text != nil {
		fields = append(fields, todo.FieldText)
	}
	if m.done != nil {
		fields = append(fields, todo.FieldDone)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *TodoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldCreatedAt:
		return m.CreatedAt()
	case todo.FieldUpdatedAt:
		return m.UpdatedAt()
	case todo.FieldText:
		return m.Text()
	case todo.FieldDone:
		return m.Done()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TodoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todo.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case todo.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case todo.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case todo.FieldDone:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDone(v)
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *TodoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *TodoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TodoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Todo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *TodoMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *TodoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *TodoMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Todo nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *TodoMutation) ResetField(name string) error {
	switch name {
	case todo.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case todo.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case todo.FieldText:
		m.ResetText()
		return nil
	case todo.FieldDone:
		m.ResetDone()
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *TodoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *TodoMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *TodoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *TodoMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *TodoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *TodoMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *TodoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Todo unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *TodoMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Todo edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	todo          map[int]struct{}
	removedtodo   map[int]struct{}
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op) *UserMutation {
	return &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on User creation.
func (m *UserMutation) SetID(id int) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddTodoIDs adds the todo edge to Todo by ids.
func (m *UserMutation) AddTodoIDs(ids ...int) {
	if m.todo == nil {
		m.todo = make(map[int]struct{})
	}
	for i := range ids {
		m.todo[ids[i]] = struct{}{}
	}
}

// RemoveTodoIDs removes the todo edge to Todo by ids.
func (m *UserMutation) RemoveTodoIDs(ids ...int) {
	if m.removedtodo == nil {
		m.removedtodo = make(map[int]struct{})
	}
	for i := range ids {
		m.removedtodo[ids[i]] = struct{}{}
	}
}

// RemovedTodo returns the removed ids of todo.
func (m *UserMutation) RemovedTodoIDs() (ids []int) {
	for id := range m.removedtodo {
		ids = append(ids, id)
	}
	return
}

// TodoIDs returns the todo ids in the mutation.
func (m *UserMutation) TodoIDs() (ids []int) {
	for id := range m.todo {
		ids = append(ids, id)
	}
	return
}

// ResetTodo reset all changes of the todo edge.
func (m *UserMutation) ResetTodo() {
	m.todo = nil
	m.removedtodo = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.todo != nil {
		edges = append(edges, user.EdgeTodo)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodo:
		ids := make([]ent.Value, 0, len(m.todo))
		for id := range m.todo {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtodo != nil {
		edges = append(edges, user.EdgeTodo)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodo:
		ids := make([]ent.Value, 0, len(m.removedtodo))
		for id := range m.removedtodo {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTodo:
		m.ResetTodo()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
