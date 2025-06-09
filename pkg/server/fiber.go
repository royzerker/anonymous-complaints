package server

import (
	userHandler "anonymous-complaints/api/v1/user"
	userRoutes "anonymous-complaints/api/v1/user"
	"anonymous-complaints/docs"
	userRepository "anonymous-complaints/internal/user"
	userService "anonymous-complaints/internal/user"

	"anonymous-complaints/pkg/logger"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"
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

	docs.SwaggerInfo.BasePath = "/" // Base path de tu API

	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	s := &Server{
		app:    app,
		log:    log,
		mongo:  mongo,
		dbName: dbName,
	}

	s.registerRoutes()

	return s
}

func (s *Server) registerRoutes() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		s.log.Info("Ping received")
		return c.SendString("Anonymous Complaints API is running!")
	})

	db := s.mongo.Database(s.dbName)
	userRepo := userRepository.NewMongoUserRepository(db)
	authUseCase := userService.NewUserService(userRepo)
	authController := userHandler.NewUserHandler(authUseCase)

	/**
	* Register routes
	 */
	userRoutes.RegisterUserRoutes(s.app, authController)
}

func (s *Server) Start(port string) {
	s.log.Info("Starting Fiber server on port " + port)
	if err := s.app.Listen(":" + port); err != nil {
		s.log.Error("Failed to start server: " + err.Error())
	}
}
