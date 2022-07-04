package out

import (
	"github.com/Adetunjii/secrot/internal/adapters/framework/out/db"
	"github.com/Adetunjii/secrot/internal/adapters/framework/out/db/entities"
)

type DatabasePort interface {
	Connect(config db.DbConfig) error
	CloseConnection() error
	RestartConnection(config db.DbConfig) error
	GetUserById(id int32) (entities.User, error)
}
