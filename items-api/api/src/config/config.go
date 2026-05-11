package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectionString = ""
	Port              = 0
)

// Load vai initializes environment variables
func Load() error {
	godotenv.Load()

	var err error
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	DatabaseConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NOME"),
	)

	return nil
}
