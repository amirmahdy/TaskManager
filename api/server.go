package api

import (
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

func NewServer(config utils.Config, store db.Store, token token.Maker) (*Server, error) {
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

	server.router = router
	return server, nil
}

func (s *Server) Run(add string) error {
	return s.router.Run(add)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
