package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseLink = ""
	Port         = 0
)

func LoadEnv() {
	var err error
	if err = godotenv.Load(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		fmt.Println(err)
		Port = 9000
	}

	DatabaseLink = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
