package users

import (
    "errors"
    "regexp"
    "strings"
)

type UserService interface {
    CreateUser(user *User) error
    GetUserByID(id int) (*User, error)
}

type service struct {
    repo UserRepository
}

func NewService(repo UserRepository) UserService {
    return &service{repo: repo}
}

func (s *service) CreateUser(user *User) error {
    if err := validateUser(user); err != nil {
        return err
    }
    return s.repo.Create(user)
}

func (s *service) GetUserByID(id int) (*User, error) {
    if id <= 0 {
        return nil, errors.New("invalid user ID")
    }
    return s.repo.GetByID(id)
}

func validateUser(user *User) error {
    if user == nil {
        return errors.New("user cannot be nil")
    }

    if user.TenantID <= 0 {
        return errors.New("tenant ID is required")
    }

    email := strings.TrimSpace(user.Email)
    if email == "" {
        return errors.New("email is required")
    }
    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    if !regexp.MustCompile(emailRegex).MatchString(email) {
        return errors.New("invalid email format")
    }

    if len(user.Name) > 255 {
        return errors.New("name exceeds maximum length of 255 characters")
    }

    return nil
}
