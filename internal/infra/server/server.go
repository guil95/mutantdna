package server

import (
	"fmt"
	"github.com/guil95/mutantdna/internal/usecase"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type server struct {
	uc *usecase.UseCase
}

func NewServer(uc *usecase.UseCase) *server {
	return &server{uc: uc}
}

func (s *server) Run() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "DNA DNASequence API"})
	})

	router.POST("/mutant", s.isMutant)
	router.GET("/stats", s.stats)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "80"), router))
}

func (s *server) isMutant(ctx *gin.Context) {
	var payload struct {
		Dna []string `json:"dna"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		handleError(err, ctx)
		return
	}

	if !validPayload(payload.Dna) {
		ctx.Status(http.StatusUnprocessableEntity)
		return
	}

	err := s.uc.SaveDna(ctx, payload.Dna)
	if err != nil {
		handleError(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *server) stats(ctx *gin.Context) {
	stats, err := s.uc.RetrieveStats(ctx)
	if err != nil {
		handleError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

func validPayload(dnas []string) bool {
	const dnaLength = 6
	pattern := "^[ATCGatcg]+$"

	if len(dnas) < dnaLength {
		return false
	}

	for _, dna := range dnas {
		if len(dna) < dnaLength {
			return false
		}

		isValid, _ := regexp.MatchString(pattern, dna)

		if isValid == false {
			return false
		}
	}

	return true
}
