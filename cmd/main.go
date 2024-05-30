package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/encountea/todo-app"
	"github.com/encountea/todo-app/pkg/handler"
	"github.com/encountea/todo-app/pkg/repository"
	"github.com/encountea/todo-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Could not initialize config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User: viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize config: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := new(todo.Server)

	log.Printf("Server is running at port: %v", viper.GetString("port"))
	if err := app.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error to connect to server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
