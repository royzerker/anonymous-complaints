package server

import (
	"anonymous-complaints/internal/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	app    *fiber.App
	log    logger.Logger
	mongo  *mongo.Client
	dbName string
}

func NewFiberServer(log logger.Logger, mongo *mongo.Client, dbName string) *Server {
	app := fiber.New()

	s := &Server{
		app:    app,
		log:    log,
		mongo:  mongo,
		dbName: dbName,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		s.log.Info("Ping received")
		return c.SendString("Anonymous Complaints API is running!")
	})

	/**
	Puedes agregar más rutas aquí
	*/
}

func (s *Server) Start(port string) {
	s.log.Info("Starting Fiber server on port " + port)
	if err := s.app.Listen(":" + port); err != nil {
		s.log.Error("Failed to start server: " + err.Error())
	}
}
