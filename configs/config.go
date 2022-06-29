package configs

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Env env
var Db *gorm.DB

type env struct {
	SecretKey string
	AppPort   int
	Debug     bool
}

func Load() error {
	Env.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	Env.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	Env.SecretKey = os.Getenv("APP_SECRET")

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), port, os.Getenv("DB_NAME"),
	)
	if Env.Debug {
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
		})
	} else {
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		})
	}

	return err
}
