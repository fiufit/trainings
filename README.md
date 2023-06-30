# trainings

Microservice for managing fiufit training plans

[![Fly Deploy](https://github.com/fiufit/trainings/actions/workflows/fly.yml/badge.svg?branch=main)](https://github.com/fiufit/trainings/actions/workflows/fly.yml)

[![test](https://github.com/fiufit/trainings/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/fiufit/trainings/actions/workflows/test.yml)
[![Lint Go Code](https://github.com/fiufit/trainings/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/fiufit/trainings/actions/workflows/lint.yml)


[![codecov](https://codecov.io/github/fiufit/trainings/branch/main/graph/badge.svg?token=CXUBV3XKVZ)](https://codecov.io/github/fiufit/trainings)

### Usage

#### With docker:
* Edit .example-env with your own secret credentials and rename it to .env
* `docker build -t fiufit-trainings .`
* `docker run -p PORT:PORT --env-file=.env fiufit-trainings`

#### Natively: 
* `go mod tidy`
* set your environvent variables to imitate the .env-example
* `go run main.go` or `go build` and run the executable


#### Running tests:
* `go test ./...`


### Links
* Fly.io deploy dashboard: https://fly.io/apps/fiufit-trainings
* Coverage report: https://app.codecov.io/github/fiufit/trainings