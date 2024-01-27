package config

import (
	"database/sql"
	"dating/constants"
	"dating/domains/entities"
	"dating/repository/sql/auth"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	AuthRepo auth.IAuthRepository
}

func InitRepository() *Repository {
	mysqlRepo := initMySql()
	return &Repository{
		AuthRepo: auth.InitAuthRepository(mysqlRepo),
	}
}

func initMySql() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=disable", constants.MySQLHost, constants.MySQLUser, constants.MySQLPassword, constants.MySQLDBName, constants.MySQLPort, constants.MySQLDBSchema)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	dborm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println("status: ", err)
		panic(err)
	}

	if constants.AutoMigrate == "true" {
		dborm.AutoMigrate(&entities.User{})
	}

	return dborm
}
