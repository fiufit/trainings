package server

import (
	"github.com/fiufit/trainings/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) InitRoutes() {
	baseRouter := s.router.Group("/:version")
	trainingRouter := baseRouter.Group("/trainings")

	s.InitTrainingRoutes(trainingRouter)

}

func (s *Server) InitTrainingRoutes(router *gin.RouterGroup) {
	router.POST("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.createTraining.Handle(),
	}))

	router.GET("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getTrainings.Handle(),
	}))

	router.PATCH("/:trainingID", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.updateTraining.Handle(),
	}))
}
