package config

import (
	"database/sql"
	"dating/constants"
	"dating/domains/entities"
	"dating/repository/sql/auth"
	"dating/repository/sql/premium_user"
	"dating/repository/sql/profile"
	"dating/repository/sql/swipe"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	AuthRepo        auth.IAuthRepository
	ProfileRepo     profile.IProfileRepository
	SwipeRepo       swipe.ISwipeRepository
	PremiumUserRepo premium_user.IPremiumUserRepository
}

func InitRepository() *Repository {
	mysqlRepo := initMySql()
	return &Repository{
		AuthRepo:        auth.InitAuthRepository(mysqlRepo),
		ProfileRepo:     profile.InitProfileRepository(mysqlRepo),
		SwipeRepo:       swipe.InitSwipeRepository(mysqlRepo),
		PremiumUserRepo: premium_user.InitPremiumUserRepository(mysqlRepo),
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
		dborm.Migrator().AutoMigrate(&entities.User{}, &entities.Swipe{}, &entities.PremiumUser{})
		//dborm.Migrator().AutoMigrate(&entities.Swipe{})
	}

	return dborm
}
