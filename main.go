package main

import (
	"log"
	"os"
	"time"

	"github.com/manolors/full-api-example/routes"

	"github.com/manolors/full-api-example/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

var DB
func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)
	DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("error al conectar a la base de datos:", err)
	}

	//	s.DB.AutoMigrate(&model.TipoCoche{}, &model.Coche{}, &model.Cliente{}, &model.Alquiler{})

	_ = routes.SetupRouter().Run()

}
