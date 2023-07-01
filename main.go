package main

import (
	"github.com/fiufit/trainings/server"
	_ "github.com/lib/pq"
)

//	@title			Fiufit Trainings API
//	@version		dev
//	@description	Fiufit's Trainings service documentation. This service manages training plans and its sessions, goals and reviews, etc.

//	@host		fiufit-trainings.fly.dev
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	srv := server.NewServer()
	srv.InitRoutes()
	srv.Run()
}
