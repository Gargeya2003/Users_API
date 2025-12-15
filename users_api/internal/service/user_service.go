package service

import (
	"context"
	"time"

	"users_api/internal/models"
	"users_api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

/* ---------------- GET USER ---------------- */

func (s *UserService) GetUser(ctx context.Context, id int32) (*models.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob,
		Age:  CalculateAge(u.Dob),
	}, nil
}

/* ---------------- CREATE USER ---------------- */

func (s *UserService) CreateUser(
	ctx context.Context,
	req models.CreateUserRequest,
) (int32, error) {

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(ctx, req.Name, dob)
}

/* ---------------- UPDATE USER ---------------- */

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	req models.UpdateUserRequest,
) error {

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, id, req.Name, dob)
}

/* ---------------- DELETE USER ---------------- */

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return s.repo.Delete(ctx, id)
}

/* ---------------- LIST ALL USERS ---------------- */

func (s *UserService) ListUsers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]models.User, error) {

	if limit <= 0 || limit > 50 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	users, err := s.repo.ListPaginated(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	result := make([]models.User, 0, len(users))
	for _, u := range users {
		result = append(result, models.User{
			ID:   int64(u.ID),
			Name: u.Name,
			DOB:  u.Dob,
			Age:  CalculateAge(u.Dob),
		})
	}

	return result, nil
}

