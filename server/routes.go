package server

import (
	"github.com/fiufit/trainings/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) InitRoutes() {
	baseRouter := s.router.Group("/:version")
	trainingRouter := baseRouter.Group("/trainings")
	exerciseRouter := trainingRouter.Group("/:trainingID/exercises")
	reviewRouter := trainingRouter.Group("/:trainingID/reviews")
	sessionRouter := baseRouter.Group("/training_sessions")
	goalsRouter := baseRouter.Group("/goals")

	s.InitTrainingRoutes(trainingRouter)
	s.InitExerciseRoutes(exerciseRouter)
	s.InitReviewRoutes(reviewRouter)
	s.InitTrainingSessionRoutes(sessionRouter)
	s.InitGoalsRoutes(goalsRouter)

}

func (s *Server) InitGoalsRoutes(router *gin.RouterGroup) {
	router.POST("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.createGoal.Handle(),
	}))
}

func (s *Server) InitTrainingSessionRoutes(router *gin.RouterGroup) {
	router.POST("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.createTrainingSession.Handle(),
	}))

	router.GET("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getTrainingSessions.Handle(),
	}))

	router.GET("/:trainingSessionID", middleware.BindTrainingSessionIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getTrainingSessionByID.Handle(),
	}))

	router.PUT("/:trainingSessionID", middleware.BindTrainingSessionIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.updateTrainingSession.Handle(),
	}))
}

func (s *Server) InitTrainingRoutes(router *gin.RouterGroup) {
	router.POST("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.createTraining.Handle(),
	}))

	router.GET("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getTrainings.Handle(),
	}))

	router.PUT("/:trainingID", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.updateTraining.Handle(),
	}))

	router.DELETE("/:trainingID", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.deleteTraining.Handle(),
	}))
}

func (s *Server) InitExerciseRoutes(router *gin.RouterGroup) {
	router.POST("", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.createExercise.Handle(),
	}))

	router.DELETE("/:exerciseID", middleware.BindTrainingIDFromUri(), middleware.BindExerciseIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.deleteExercise.Handle(),
	}))

	router.PATCH("/:exerciseID", middleware.BindTrainingIDFromUri(), middleware.BindExerciseIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.updateExercise.Handle(),
	}))

	router.GET("/:exerciseID", middleware.BindTrainingIDFromUri(), middleware.BindExerciseIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getExercise.Handle(),
	}))
}

func (s *Server) InitReviewRoutes(router *gin.RouterGroup) {
	router.POST("", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.createReview.Handle(),
	}))

	router.GET("", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getReviews.Handle(),
	}))

	router.GET("/:reviewID", middleware.BindTrainingIDFromUri(), middleware.BindReviewIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getReviewByID.Handle(),
	}))

	router.PATCH("/:reviewID", middleware.BindTrainingIDFromUri(), middleware.BindReviewIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.updateReview.Handle(),
	}))

	router.DELETE("/:reviewID", middleware.BindTrainingIDFromUri(), middleware.BindReviewIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.deleteReview.Handle(),
	}))

}
