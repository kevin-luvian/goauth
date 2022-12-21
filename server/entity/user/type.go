package user

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Tag       string    `json:"tag" db:"tag"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	HPass     string    `json:"-" db:"hpass"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
