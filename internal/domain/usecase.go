package domain

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"strings"
)

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) SaveDna(ctx context.Context, dna DNA) error {
	isMutant := IsMutant(dna)

	hashDna := md5.Sum([]byte(strings.Join(dna, "")))

	mutant := Mutant{ID: fmt.Sprintf("%x", hashDna), DNA: dna, Type: HumanDNA}
	if isMutant {
		mutant.Type = MutantDNA
	}

	go func() {
		err := u.repo.SaveDna(ctx, mutant)
		if err != nil {
			log.Println(fmt.Sprintf("error in save mutant %s", err.Error()))
		}
	}()

	if mutant.Type == MutantDNA {
		return nil
	}

	return HumanError{}
}

func (u *UseCase) RetrieveStats(ctx context.Context) (*Stats, error){
	return u.repo.FindStats(ctx)
}
