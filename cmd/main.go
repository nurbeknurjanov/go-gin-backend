package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	go_backend "github.com/nurbeknurjanov/go-gin-backend"
	"github.com/nurbeknurjanov/go-gin-backend/grpc"
	grpc_handlers "github.com/nurbeknurjanov/go-gin-backend/grpc/handlers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/handlers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
)

// @title CRUD App API
// @version 1.0
// @description API Server for CRUD Application

// @host localhost:3001
// @BasePath /

// @securityDefinitions.apiKey AccessTokenHeaderName
// @in header
// @name X-Access-Token

func main() {
	//logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error reading configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error reading env variables: %s", err.Error())
	}

	db, err := repositories.NewPostgresDb(repositories.DbConfig{
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

	repo := repositories.NewSqlRepositories(db)

	/*producer, err := k.NewProducer([]string{os.Getenv("KAFKA1_HOST"), os.Getenv("KAFKA2_HOST"), os.Getenv("KAFKA3_HOST")})
	if err != nil {
		logrus.Fatalf("error connecting to kafka nodes : %s", err.Error())
	}*/

	s := services.NewServices(repo, nil)

	handler := handlers.NewHandler(s)

	server := new(go_backend.Server)

	grpcHandlers := grpc_handlers.NewGrpcHandlers(grpc_handlers.Deps{
		Auth: s.Auth,
	})
	grpcServer := grpc.NewServer(grpc.Deps{
		AuthHandler: grpcHandlers.AuthHandler,
	})
	go func() {
		if err := grpcServer.ListenAndServer(3002); err != nil {
			logrus.Fatalf("error starting grpc server: %s", err.Error())
		}
	}()

	if err := server.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error starting server: %s", err.Error())
	}
}

func initConfig() error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	viper.SetConfigFile(dir + "/../configs/config.yaml")
	return viper.ReadInConfig()
}
