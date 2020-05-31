package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/art-Hasan/gqlgen-todos/ent"
	"github.com/art-Hasan/gqlgen-todos/ent/todo"
	"github.com/art-Hasan/gqlgen-todos/ent/user"
	"github.com/art-Hasan/gqlgen-todos/graph/generated"
	"github.com/art-Hasan/gqlgen-todos/graph/model"
	"github.com/art-Hasan/gqlgen-todos/pkg/db"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*ent.Todo, error) {
	tx, err := db.Start(ctx, r.ent)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	t, err := tx.Todo.Create().
		SetText(input.Text).
		SetUserID(input.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("createTodo: db error: %w", err)
	}

	return t, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*ent.User, error) {
	tx, err := db.Start(ctx, r.ent)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	u, err := tx.User.Create().
		SetName(input.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("createUser: db error: %w", err)
	}
	return u, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*ent.Todo, error) {
	tx, err := db.Start(ctx, r.ent)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if _, err := tx.Todo.Delete().
		Where(todo.IDEQ(input.ID)).
		Exec(ctx); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("deleteTodo: db error: %w", err)
	}
	return nil, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUser) (*ent.User, error) {
	tx, err := db.Start(ctx, r.ent)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if _, err := tx.User.Delete().
		Where(user.IDEQ(input.ID)).
		Exec(ctx); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("deleteUser: db error: %w", err)
	}
	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) (*model.TodosPage, error) {
	q := r.ent.Todo.Query()
	all, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	total, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &model.TodosPage{
		Page: all,
		Info: &model.PageInfo{Total: total},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) (*model.UsersPage, error) {
	q := r.ent.User.Query()
	all, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	total, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &model.UsersPage{
		Page: all,
		Info: &model.PageInfo{Total: total},
	}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *ent.Todo) (*ent.User, error) {
	return r.ent.Todo.QueryUser(obj).Only(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
