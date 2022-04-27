package usecase

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/guil95/mutantdna/internal/domain"
	"log"
	"strings"
)

type UseCase struct {
	repo domain.Repository
}

func NewUseCase(repo domain.Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) SaveDna(ctx context.Context, dnaSequence domain.DNASequence) error {
	hashDna := md5.Sum([]byte(strings.Join(dnaSequence, "")))

	dna := domain.DNA{ID: fmt.Sprintf("%x", hashDna), DNASequence: dnaSequence, Type: domain.HumanDNA}

	isMutant := domain.IsMutant(dnaSequence)
	if isMutant {
		dna.Type = domain.MutantDNA
	}

	go func() {
		err := u.repo.SaveDna(ctx, dna)
		if err != nil {
			log.Println(fmt.Sprintf("error in save mutant %s", err.Error()))
		}
	}()

	if dna.Type == domain.MutantDNA {
		return nil
	}

	return domain.HumanError{}
}

func (u *UseCase) RetrieveStats(ctx context.Context) (*domain.Stats, error){
	stats, err := u.repo.FindStats(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return stats, nil
}
