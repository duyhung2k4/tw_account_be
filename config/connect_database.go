package config

import (
	"account-service/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect(migrate bool) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		`	host=%s
			user=%s
			password=%s
			dbname=%s
			port=%s
			sslmode=disable`,
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	var err error

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Connected")

	if migrate {
		log.Println("Start migrate")
		err := db.AutoMigrate(
			&model.Profile{},
			&model.Credential{},
			&model.Role{},
			&model.SaveRegister{},
		)

		if err != nil {
			return nil, err
		}

		log.Println("Migrate done!")
	}

	return db, nil
}
