package main

import (
	"github.com/guil95/mutantdna/config"
	"github.com/guil95/mutantdna/internal/domain"
	"github.com/guil95/mutantdna/internal/infra/repositories"
	"github.com/guil95/mutantdna/internal/infra/server"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	repo := repositories.NewRepository(config.GetDBConnection())
	uc := domain.NewUseCase(repo)
	s := server.NewServer(uc)

	s.Run()
}
