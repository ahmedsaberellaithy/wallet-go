package users

import "time"

type User struct {
    ID        int       `gorm:"primaryKey"`
    TenantID  int       `gorm:"not null"`
    Name      string    `gorm:"size:255"`
    Email     string    `gorm:"uniqueIndex"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
