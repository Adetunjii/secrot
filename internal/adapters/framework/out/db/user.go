package db

import (
	. "github.com/Adetunjii/secrot/internal/adapters/framework/out/db/entities"
)

func (db *Database) GetUserById(userId int32) (User, error) {
	var user User
	err := db.connection.First(&user, userId).Error
	if err != nil {
		// error handling logic
	}

	return user, nil
}
