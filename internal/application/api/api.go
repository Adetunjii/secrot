package api

import (
	. "github.com/Adetunjii/secrot/internal/adapters/framework/out/db/entities"
	ports "github.com/Adetunjii/secrot/internal/ports/out"
)

type Application struct {
	db   ports.DatabasePort
	core Core
}

func NewApplication(db ports.DatabasePort, core Core) *Application {
	return &Application{db: db, core: core}
}

func (api *Application) Deposit(userId int32, amount int32) (User, error) {
	user, err := api.db.GetUserById(userId)
	if err != nil {
		// error handling logic
		return User{}, err
	}

	result := api.core.Deposit(user, amount)
	return result, nil
}
