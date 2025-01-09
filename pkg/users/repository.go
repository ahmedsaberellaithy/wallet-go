package users

import "wallet-sdk/internal/database"

type UserRepository interface {
    Create(user *User) error
    GetByID(id int) (*User, error)
}

type repository struct {
    db database.DBInterface
}

func NewRepository(db database.DBInterface) UserRepository {
    return &repository{db: db}
}

func (r *repository) Create(user *User) error {
    result := r.db.Create(user)
    return result.Error
}

func (r *repository) GetByID(id int) (*User, error) {
    var user User
    result := r.db.First(&user, id)
    return &user, result.Error
}
