package constants

import "os"

var (
	Port                 = ""
	MySQLHost     string = os.Getenv("DB_HOST")
	MySQLUser     string = os.Getenv("DB_USER")
	MySQLPassword string = os.Getenv("DB_PASSWORD")
	MySQLPort     string = os.Getenv("DB_PORT")
	MySQLDBName   string = os.Getenv("DB_NAME")
	MySQLDBSchema string = os.Getenv("DB_SCHEMA")
	AutoMigrate   string = ""
)
