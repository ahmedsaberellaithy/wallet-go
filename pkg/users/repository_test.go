package users

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "gorm.io/gorm"
)

type mockDB struct {
    mock.Mock
}

func (m *mockDB) Create(value interface{}) *gorm.DB {
    args := m.Called(value)
    return args.Get(0).(*gorm.DB)
}

func (m *mockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
    args := m.Called(dest, conds)
    return args.Get(0).(*gorm.DB)
}

func TestRepository_Create(t *testing.T) {
    mockDB := new(mockDB)
    repo := NewRepository(mockDB)

    user := &User{
        TenantID:  123,
        Name:      "John Doe",
        Email:     "john.doe@example.com",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    mockDB.On("Create", user).Return(&gorm.DB{Error: nil})

    err := repo.Create(user)
    assert.NoError(t, err)
    mockDB.AssertExpectations(t)
}

func TestRepository_GetByID(t *testing.T) {
    mockDB := new(mockDB)
    repo := NewRepository(mockDB)

    expectedUser := &User{
        ID:        1,
        TenantID:  123,
        Name:      "John Doe",
        Email:     "john.doe@example.com",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    mockDB.On("First", mock.AnythingOfType("*users.User"), mock.Anything).Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*User)
        *arg = *expectedUser
    })

    user, err := repo.GetByID(1)
    assert.NoError(t, err)
    assert.Equal(t, expectedUser, user)
    mockDB.AssertExpectations(t)
}
