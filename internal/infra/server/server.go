package server

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/guil95/mutantdna/internal/domain"
)

type server struct {
	uc *domain.UseCase
}

func NewServer(uc *domain.UseCase) *server {
	return &server{uc: uc}
}

func (s *server) Run() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Mutant DNA API"})
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
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if !validPayload(payload.Dna) {
		ctx.Status(http.StatusUnprocessableEntity)
		return
	}

	err := s.uc.SaveDna(ctx, payload.Dna)
	if err != nil {
		ctx.Status(http.StatusForbidden)
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *server) stats(ctx *gin.Context) {
	stats, err := s.uc.RetrieveStats(ctx)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if stats == nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

func validPayload(dnas []string) bool {
	const dnaLength = 6
	pattern := "^[ATCGatcg]+$"

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
