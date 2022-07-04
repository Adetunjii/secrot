package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DbConfig struct {
	Host        string
	Port        string
	User        string
	Password    string
	Name        string
	DatabaseUrl string
}

type Database struct {
	connection *gorm.DB
}

func NewAdapter(config DbConfig) (*Database, error) {
	db := &Database{
		connection: nil,
	}

	if err := db.Connect(config); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) Connect(config DbConfig) error {

	// using gorm to connect to postgres driver
	var dsn string
	databaseUrl := config.DatabaseUrl
	if databaseUrl == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Name)

	} else {
		dsn = databaseUrl
	}

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.connection = connection
	log.Println("Database Connected Successfully...")

	return nil
}

func (db *Database) CloseConnection() error {
	if db.connection != nil {
		conn, err := db.connection.DB()
		if err != nil {
			return err
		}

		conn.Close()
	}
	return nil
}

func (db *Database) RestartConnection(config DbConfig) error {

	// close existing connections if any
	if db.connection != nil {
		db.CloseConnection()
	}

	// reconnect and swap out the old connection
	return db.Connect(config)
}
