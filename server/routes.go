package server

import (
	_ "github.com/fiufit/trainings/docs"
	"github.com/fiufit/trainings/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	s.InitDocRoutes(baseRouter)

}

func (s *Server) InitDocRoutes(router *gin.RouterGroup) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) InitGoalsRoutes(router *gin.RouterGroup) {
	router.POST("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.createGoal.Handle(),
	}))

	router.GET("", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getGoals.Handle(),
	}))

	router.GET("/:goalID", middleware.BindGoalIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getGoalByID.Handle(),
	}))

	router.PATCH("/:goalID", middleware.BindGoalIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.updateGoal.Handle(),
	}))

	router.DELETE("/:goalID", middleware.BindGoalIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.deleteGoal.Handle(),
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

	router.POST("/:trainingID/favorites", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.addFavorite.Handle(),
	}))

	router.DELETE("/:trainingID/favorites", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.removeFavorite.Handle(),
	}))

	router.GET("/favorites", middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.getFavorites.Handle(),
	}))

	router.POST("/:trainingID/enable", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.enableTraining.Handle(),
	}))

	router.DELETE("/:trainingID/disable", middleware.BindTrainingIDFromUri(), middleware.HandleByVersion(middleware.VersionHandlers{
		"v1": s.disableTraining.Handle(),
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
