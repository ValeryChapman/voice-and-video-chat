package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	chat "github.com/ValeryChapman/chat"
	"github.com/ValeryChapman/chat/pkg/handler"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	// database
	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//})
	//if err != nil {
	//	logrus.Fatalf("failed to initialize db: %s", err.Error())
	//}

	//repos := repository.NewRepository(db)
	//services := service.NewService(repos)
	handlers := handler.NewHandler()

	srv := new(chat.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Voice And Video Chat App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Voice And Video Chat App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	//if err := db.Close(); err != nil {
	//	logrus.Errorf("error occured on db connection close: %s", err.Error())
	//}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
