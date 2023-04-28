package main

import "github.com/fiufit/trainings/server"

func main() {
	srv := server.NewServer()
	srv.InitRoutes()
	srv.Run()
}
