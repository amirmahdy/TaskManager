package api

import (
	"log"
	db "taskmanager/db/model"
	"taskmanager/docs"
	"taskmanager/token"
	"taskmanager/utils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config utils.Config
	token  token.Maker
	store  db.Store
	router *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	token, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker", err)
	}
	server := &Server{
		config: config,
		store:  store,
		token:  token,
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.POST("/api/user/login", server.loginUser)
	router.POST("/api/user/create", server.createUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.token))
	authRoutes.GET("/api/task", server.getTask)
	authRoutes.POST("/api/task", server.createTask)

	server.router = router
	return server, nil
}

func (s *Server) Run(add string) error {
	return s.router.Run(add)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
