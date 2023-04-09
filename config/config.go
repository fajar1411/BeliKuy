package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"fmt"

)

type AppConfig struct {
	DB_USERNAME              string
	DB_PASSWORD              string
	DB_HOST                  string
	DB_PORT                  uint
	DB_NAME                  string
	SERVER_PORT              uint
	JWT_SECRET               string
	CLOUDINARY_CLOUD_NAME    string
	CLOUDINARY_API_KEY       string
	CLOUDINARY_API_SECRET    string
	CLOUDINARY_UPLOAD_FOLDER string
	ENV_POST string
	COLLECTION_POST string
	POST_KEY string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	// if _, exist := os.LookupEnv("SECRET"); !exist {
	// 	if err := godotenv.Load("local.env"); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// SECRET = os.Getenv("SECRET")
	cnvServerPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	fmt.Println("ini port", cnvServerPort)
	if err != nil {
		log.Fatal("Cannot parse Server Port variable")
		return nil
	}
	defaultConfig.SERVER_PORT = uint(cnvServerPort)
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	defaultConfig.CLOUDINARY_API_KEY = os.Getenv("CLOUDINARY_API_KEY")
	defaultConfig.CLOUDINARY_API_SECRET = os.Getenv("CLOUDINARY_API_SECRET")
	defaultConfig.CLOUDINARY_CLOUD_NAME = os.Getenv("CLOUDINARY_CLOUD_NAME")
	defaultConfig.CLOUDINARY_UPLOAD_FOLDER = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	cnvDBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.DB_PORT = uint(cnvDBPort)
	defaultConfig.JWT_SECRET = os.Getenv("JWT_SECRET")
	defaultConfig.ENV_POST = os.Getenv("ENV_POST")
	defaultConfig.COLLECTION_POST = os.Getenv("COLLECTION_POST")
	defaultConfig.POST_KEY = os.Getenv("POST_KEY")

	return &defaultConfig
}
