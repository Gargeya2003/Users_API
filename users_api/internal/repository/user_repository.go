package repository

import (
	"context"
	"time"

	"users_api/db/sqlc"
)

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q}
}

// CREATE
func (r *UserRepository) Create(
	ctx context.Context,
	name string,
	dob time.Time,
) (int32, error) {
	res, err := r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int32(id), err
}

// GET BY ID
func (r *UserRepository) GetByID(
	ctx context.Context,
	id int32,
) (*sqlc.User, error) {
	u, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// LIST
func (r *UserRepository) List(
	ctx context.Context,
) ([]sqlc.User, error) {
	return r.q.ListUsers(ctx)
}

// UPDATE
func (r *UserRepository) Update(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) error {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

// DELETE
func (r *UserRepository) Delete(
	ctx context.Context,
	id int32,
) error {
	return r.q.DeleteUser(ctx, id)
}

// LIST PAGINATED
func (r *UserRepository) ListPaginated(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]sqlc.User, error) {
	return r.q.ListUsersPaginated(ctx, sqlc.ListUsersPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
}
