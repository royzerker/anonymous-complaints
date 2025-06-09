package main

import (
	"anonymous-complaints/pkg/config"
	"anonymous-complaints/pkg/db/mongodb"
	"anonymous-complaints/pkg/logger"
	"anonymous-complaints/pkg/server"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found or failed to load.")
	}
}

// @title API de Quejas Anónimas
// @version 1.0
// @description Esta es la API para manejar usuarios y quejas anónimas.
// @host localhost:8000
// @BasePath /
func main() {
	cfg := config.ConfigService()

	logLevel := cfg.GetAndCheck("LOG_LEVEL")

	var level logger.LogLevel
	switch logLevel {
	case "DEBUG":
		level = logger.DEBUG
	case "WARN":
		level = logger.WARN
	case "ERROR":
		level = logger.ERROR
	default:
		level = logger.INFO
	}
	logg := logger.RequestLogger(level)

	mongoURI := cfg.GetAndCheck("MONGO_URI")

	mongoClient, err := mongodb.MongoClient(mongoURI)

	if err != nil {
		logg.Error("Failed to connect to MongoDB: " + err.Error())
		return
	}

	logg.Info("Connected to MongoDB successfully")

	dbName := "agnostic"

	srv := server.NewFiberServer(logg, mongoClient, dbName)
	port := cfg.GetAndCheck("SERVER_PORT")
	srv.Start(port)
}
