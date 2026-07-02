package main

import (
	"parkora/internal/config"
	"parkora/internal/database"
	"parkora/internal/server"
)

func main() {
	envs := config.LoadEnv()
	db := database.Pdb(envs)
	server.StartServer(db)
}
