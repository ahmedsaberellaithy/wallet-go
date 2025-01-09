package users

import (
    "testing"
    "time"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockRepository struct {
    mock.Mock
}

func (m *MockRepository) Create(user *User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockRepository) GetByID(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

func TestService_CreateUser(t *testing.T) {
    mockRepo := new(MockRepository)
    service := NewService(mockRepo)

    tests := []struct {
        name    string
        user    *User
        mockFn  func()
        wantErr bool
        errMsg  string
    }{
        {
            name: "successful creation",
            user: &User{
                TenantID:  123,
                Name:      "John Doe",
                Email:     "john.doe@example.com",
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            },
            mockFn: func() {
                mockRepo.On("Create", mock.AnythingOfType("*users.User")).Return(nil).Once()
            },
            wantErr: false,
        },
        {
            name: "invalid tenant ID",
            user: &User{
                TenantID:  0,
                Name:      "John Doe",
                Email:     "john.doe@example.com",
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            },
            mockFn:  func() {},
            wantErr: true,
            errMsg:  "tenant ID is required",
        },
        {
            name: "invalid email format",
            user: &User{
                TenantID:  123,
                Name:      "John Doe",
                Email:     "invalid-email",
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            },
            mockFn:  func() {},
            wantErr: true,
            errMsg:  "invalid email format",
        },
        {
            name: "missing tenant ID",
            user: &User{
                Name:      "John Doe",
                Email:     "john.doe@example.com",
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            },
            mockFn:  func() {},
            wantErr: true,
            errMsg:  "tenant ID is required",
        },
        {
            name: "missing email",
            user: &User{
                TenantID:  123,
                Name:      "John Doe",
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            },
            mockFn:  func() {},
            wantErr: true,
            errMsg:  "email is required",
        },
        {
            name:    "nil user",
            user:    nil,
            mockFn:  func() {},
            wantErr: true,
            errMsg:  "user cannot be nil",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.mockFn()
            err := service.CreateUser(tt.user)
            if tt.wantErr {
                assert.Error(t, err)
                if tt.errMsg != "" {
                    assert.Equal(t, tt.errMsg, err.Error())
                }
            } else {
                assert.NoError(t, err)
            }
            mockRepo.AssertExpectations(t)
        })
    }
}

func TestService_GetUserByID(t *testing.T) {
    mockRepo := new(MockRepository)
    service := NewService(mockRepo)

    tests := []struct {
        name    string
        id      int
        want    *User
        mockFn  func()
        wantErr bool
        errMsg  string
    }{
        {
            name: "successful retrieval",
            id:   1,
            want: &User{
                ID:       1,
                TenantID: 123,
                Name:     "John Doe",
                Email:    "john.doe@example.com",
            },
            mockFn: func() {
                mockRepo.On("GetByID", 1).Return(&User{
                    ID:       1,
                    TenantID: 123,
                    Name:     "John Doe",
                    Email:    "john.doe@example.com",
                }, nil).Once()
            },
            wantErr: false,
        },
        {
            name:    "invalid ID",
            id:      0,
            want:    nil,
            mockFn:  func() {},
            wantErr: true,
            errMsg:  "invalid user ID",
        },
        {
            name:    "negative ID",
            id:      -1,
            want:    nil,
            mockFn:  func() {},
            wantErr: true,
            errMsg:  "invalid user ID",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.mockFn()
            got, err := service.GetUserByID(tt.id)
            if tt.wantErr {
                assert.Error(t, err)
                if tt.errMsg != "" {
                    assert.Equal(t, tt.errMsg, err.Error())
                }
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.want, got)
            }
            mockRepo.AssertExpectations(t)
        })
    }
}
