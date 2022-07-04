package ports

import . "github.com/Adetunjii/secrot/internal/adapters/framework/out/db/entities"

type APIPort interface {
	Deposit(userId int32, amount int32) (User, error)
}
