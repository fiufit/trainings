package repositories

import (
	"os"
	"testing"

	"github.com/fiufit/trainings/models"
	testingUtils "github.com/fiufit/trainings/utils/testing"
)

var testSuite testingUtils.TestSuite

func TestMain(m *testing.M) {
	testSuite = testingUtils.NewTestSuite(
		models.TrainingPlan{},
		&models.Exercise{},
		&models.Review{},
		&models.Tag{},
		&models.TrainingSession{},
		&models.ExerciseSession{},
		&models.Goal{},
		&models.Favorite{},
	)

	testResult := m.Run()
	testSuite.TearDown()
	os.Exit(testResult)
}
