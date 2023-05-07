package usecases

// func TestCreateTrainingOk(t *testing.T) {

// 	ctx := context.Background()
// 	creationDate := time.Now()
// 	_, err := mpatch.PatchMethod(time.Now, func() time.Time {
// 		return creationDate
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req := training.CreateTrainingRequest{
// 		Name:        "Test Name",
// 		Description: "Test Description",
// 		TrainerID:   "Test Trainer",
// 		Exercises:   []training.ExerciseRequest{},
// 	}
// 	trainingRepo := new(mocks.TrainingPlans)

// 	training := training.ConverToTrainingPlan(req)
// 	trainingRepo.On("CreateTrainingPlan", ctx, training).Return(training, nil)

// 	trainingUc := NewTrainingCreatorImpl(trainingRepo, zaptest.NewLogger(t))
// 	res, err := trainingUc.CreateTraining(ctx, req)

// 	assert.NoError(t, err)
// 	assert.Equal(t, res.TrainingPlan, training)
// }

// func TestCreateTrainingError(t *testing.T) {

// 	ctx := context.Background()
// 	req := training.CreateTrainingRequest{
// 		Name:        "Test Name",
// 		Description: "Test Description",
// 		TrainerID:   "Test Trainer",
// 		Exercises:   []training.ExerciseRequest{},
// 	}
// 	trainingRepo := new(mocks.TrainingPlans)

// 	training := training.ConverToTrainingPlan(req)
// 	trainingRepo.On("CreateTrainingPlan", ctx, training).Return(models.TrainingPlan{}, errors.New("repo error"))

// 	trainingUc := NewTrainingCreatorImpl(trainingRepo, zaptest.NewLogger(t))
// 	res, err := trainingUc.CreateTraining(ctx, req)

// 	assert.Equal(t, res.TrainingPlan, models.TrainingPlan{})
// 	assert.Error(t, err)
// }
