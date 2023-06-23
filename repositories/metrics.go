package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	mContracts "github.com/fiufit/trainings/contracts/metrics"
	"github.com/fiufit/trainings/utils"
	"go.uber.org/zap"
)

//go:generate mockery --name Metrics
type Metrics interface {
	Create(ctx context.Context, req mContracts.CreateMetricRequest)
}

type MetricsRepository struct {
	metricsUrl string
	version    string
	logger     *zap.Logger
}

func NewMetricsRepository(metricsUrl string, version string, logger *zap.Logger) MetricsRepository {
	return MetricsRepository{metricsUrl: metricsUrl, version: version, logger: logger}
}

func (repo MetricsRepository) Create(ctx context.Context, req mContracts.CreateMetricRequest) {
	url := repo.metricsUrl + "/" + repo.version + "/metrics"

	reqBytes, err := json.Marshal(req)
	if err != nil {
		repo.logger.Error("Unable to marshal metric struct before send")
		return
	}

	res, err := utils.MakeRequest(http.MethodPost, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		repo.logger.Error("Unable to make CreateMetric request")
		return
	}

	defer res.Body.Close()
	if res.StatusCode >= 400 {
		repo.logger.Error("Metrics service was unable to create metric")
	}
}
