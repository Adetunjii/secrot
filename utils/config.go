package utils

import (
	"fmt"
	"github.com/Adetunjii/secrot/internal/adapters/framework/out/db"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
	"time"
)

type Config struct {
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT""`
	DbUser       string `mapstructure:"DB_USER"`
	DbName       string `mapstructure:"DB_NAME"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`
	DbUrl        string `mapstructure:"DB_URL"`
	RabbitMQHost string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort int    `mapstructure:"RABBITMQ_PORT"`
	RabbitMQUser string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPass string `mapstructure:"RABBITMQ_PASS"`
	CloudAMQPUrl string `mapstructure:"CLOUDAMQP_URL"`
}

type Service struct {
	DB *db.Database
}

// LoadConfig: use viper to load all remote configurations
func LoadConfig(path string) (*Service, error) {

	config := &Config{}
	service := &Service{}

	consulUrl := os.Getenv("CONSUL_URL")
	consulKey := os.Getenv("CONSUL_KEY")

	fmt.Println(consulUrl, consulKey)

	remoteViper := viper.New()

	remoteViper.AddRemoteProvider("consul", consulUrl, consulKey)
	remoteViper.SetConfigType("json")

	err := remoteViper.ReadRemoteConfig()
	if err != nil {
		return &Service{}, err
	}

	err = remoteViper.Unmarshal(&config)
	if err != nil {
		return &Service{}, err
	}

	dbConfig := db.DbConfig{
		Host:     config.DbHost,
		Port:     config.DbPort,
		User:     config.DbUser,
		Password: config.DbPassword,
		Name:     config.DbName,
	}

	database, err := db.NewAdapter(dbConfig)
	if err != nil {
		return &Service{}, err
	}

	service.DB = database

	// start watcher in another goroutine that checks for an update every 30 seconds
	go func() {
		for {
			time.Sleep(30 * time.Second)

			if err := remoteViper.WatchRemoteConfig(); err != nil {
				log.Fatalf("couldn't fetch new credentials: %v", err)
			}

			remoteViper.Unmarshal(&config)

			newDbConfig := db.DbConfig{
				Host:     config.DbHost,
				Port:     config.DbPort,
				User:     config.DbUser,
				Password: config.DbPassword,
				Name:     config.DbName,
			}

			if newDbConfig != dbConfig {
				service.DB.RestartConnection(newDbConfig)
				dbConfig = newDbConfig
			}
		}
	}()

	return service, nil
}
