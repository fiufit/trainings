package contracts

var errCodes = map[error]string{
	ErrInternal:                       "T0",
	ErrBadRequest:                     "T1",
	ErrTrainingPlanNotFound:           "T2",
	ErrExerciseNotFound:               "T3",
	ErrUnauthorizedTrainer:            "T4",
	ErrSelfReview:                     "T5",
	ErrReviewAlreadyExists:            "T6",
	ErrReviewNotFound:                 "T7",
	ErrUnauthorizedReviewer:           "T8",
	ErrInvalidTag:                     "T9",
	ErrTrainingSessionNotFound:        "T10",
	ErrTrainingSessionNotComplete:     "T11",
	ErrTrainingSessionAlreadyFinished: "T12",
	ErrUserInternal:                   "U0",
	ErrUserBadRequest:                 "U1",
	ErrUserNotFound:                   "U2",
}

var externalCodes = map[string]error{
	"U0": ErrUserInternal,
	"U1": ErrUserBadRequest,
	"U2": ErrUserNotFound,
}

type OkResponse struct {
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	Err ErrPayload `json:"error"`
}

type ErrPayload struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func FormatOkResponse(data interface{}) OkResponse {
	return OkResponse{data}
}

func FormatErrResponse(err error) ErrResponse {
	errCode, ok := errCodes[err]
	if !ok {
		errCode = "T0"
	}

	payload := ErrPayload{
		Description: err.Error(),
		Code:        errCode,
	}

	return ErrResponse{payload}
}
