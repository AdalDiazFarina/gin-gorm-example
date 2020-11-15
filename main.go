package main

import (
	"log"
	"os"
	"time"

	"github.com/manolors/full-api-example/model"
	"github.com/manolors/full-api-example/routes"
	s "github.com/manolors/full-api-example/storage"

	"github.com/manolors/full-api-example/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)
	s.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("error al conectar a la base de datos:", err)
	}

	s.DB.AutoMigrate(&model.TipoCoche{}, &model.Coche{}, &model.Cliente{}, &model.Alquiler{})

	routes.SetupRouter().Run()

}
