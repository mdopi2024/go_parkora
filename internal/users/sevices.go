package users

import (
	"errors"

	"parkora/internal/auth"
	httpresponse "parkora/internal/httpResponse"
	userdto "parkora/internal/users/dto"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserService struct {
	repo Repository
	auth auth.JWTService
}

func NewUserService(repo Repository, auth auth.JWTService) *UserService {
	return &UserService{repo: repo, auth: auth}
}

func (s *UserService) Register(req userdto.RegisterUserRequest) (userdto.RegisterUserResponse, *httpresponse.ErrorResponse) {
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return userdto.RegisterUserResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to hash password",
			Errors:  err,
		}
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	if err := s.repo.CreateUser(user); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "idx_users_email" {
				return userdto.RegisterUserResponse{}, &httpresponse.ErrorResponse{
					Success: false,
					Message: "Email already exists",
					Errors: map[string]string{
						"email": "Email already exists",
					},
				}
			}
		}

		return userdto.RegisterUserResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Errors:  err.Error(),
		}
	}

	return user.ToCreateResponse("User registered successfully"), nil
}

func (s *UserService) Login(req userdto.LoginUserRequest) (userdto.LoginUserResponse, *httpresponse.ErrorResponse) {
	if req.Email == "" {
		return userdto.LoginUserResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "email is required",
			Errors:  errors.New("email is required"),
		}
	}
	if req.Password == "" {
		return userdto.LoginUserResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "password is required",
			Errors:  errors.New("password is required"),
		}
	}

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return userdto.LoginUserResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "invalid credentials",
			Errors:  errors.New("invalid credentials"),
		}
	}

	if !auth.CheckPassword(user.Password, req.Password) {
		return userdto.LoginUserResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "invalid credentials",
			Errors:  errors.New("invalid credentials"),
		}
	}

	token, err := s.auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return userdto.LoginUserResponse{}, &httpresponse.ErrorResponse{
			Success: false,
			Message: "failed to generate token",
			Errors:  errors.New("failed to generate token"),
		}
	}

	return user.ToLoginResponse(token, "Login successful"), nil
}

func (s *UserService) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) FindByEmail(email string) (*User, error) {
	return s.repo.FindByEmail(email)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
