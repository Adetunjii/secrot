package api

import . "github.com/Adetunjii/secrot/internal/adapters/framework/out/db/entities"

type Core interface {
	Deposit(user User, amount int32) User
}
