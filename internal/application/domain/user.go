package domain

import (
	. "github.com/Adetunjii/secrot/internal/adapters/framework/out/db/entities"
)

type Domain struct{}

func New() *Domain {
	return &Domain{}
}

func (d *Domain) Deposit(user User, amount int32) User {
	user.Balance += amount
	return user
}
