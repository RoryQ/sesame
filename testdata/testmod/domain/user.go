package domain

import "time"

type User struct {
	ID        string
	Name      string
	UpdatedAt time.Time
	DeletedAt *time.Time
	Address   *Address
}

func (u User) GetName() string {
	return u.Name
}
