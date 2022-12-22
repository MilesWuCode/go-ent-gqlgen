// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-ent-gqlgen/ent/todo"
	"time"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Text      string
	Status    *todo.Status
	CreatedAt *time.Time
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetText(i.Text)
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	Text   *string
	Status *todo.Status
}

// Mutate applies the UpdateTodoInput on the TodoMutation builder.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdate builder.
func (c *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdateOne builder.
func (c *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
