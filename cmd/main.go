package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	go_backend "github.com/nurbeknurjanov/go-gin-backend"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/handler"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repository"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	helpers.RunPanic()
	return
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error reading configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error reading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.DbConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		//Password: viper.GetString("db.password"),
		DbName:  viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error connecting to database: %s", err.Error())
	}

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)
	handlers := handler.NewHandler(services)

	server := new(go_backend.Server)
	if err := server.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error starting server: %s", err.Error())
	}
}

func initConfig() error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	viper.SetConfigFile(dir + "/../configs/config.yaml")
	return viper.ReadInConfig()
}
