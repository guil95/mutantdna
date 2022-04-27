package usecase

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/guil95/mutantdna/internal/domain"
	"github.com/guil95/mutantdna/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	saveDna   = "SaveDna"
	findStats = "FindStats"
)

func TestUseCase(t *testing.T) {
	t.Run("create mutant dna with success", func(t *testing.T) {
		repo := new(mocks.Repository)
		ctx := context.Background()
		mutantDnaSequence := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

		hashDna := md5.Sum([]byte(strings.Join(mutantDnaSequence, "")))

		mutantDNA := domain.DNA{ID: fmt.Sprintf("%x", hashDna), DNASequence: mutantDnaSequence, Type: domain.MutantDNA}
		repo.On(saveDna, ctx, mutantDNA).Return(nil)

		uc := NewUseCase(repo)
		err := uc.SaveDna(ctx, mutantDnaSequence)

		assert.NoError(t, err)
	})

	t.Run("create human dna should return error", func(t *testing.T) {
		repo := new(mocks.Repository)
		ctx := context.Background()
		humanDnaSequence := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAGAG", "CCCTTA", "TCACTG"}

		hashDna := md5.Sum([]byte(strings.Join(humanDnaSequence, "")))

		mutantDNA := domain.DNA{ID: fmt.Sprintf("%x", hashDna), DNASequence: humanDnaSequence, Type: domain.HumanDNA}
		repo.On(saveDna, ctx, mutantDNA).Return(nil)

		uc := NewUseCase(repo)
		err := uc.SaveDna(ctx, humanDnaSequence)

		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.HumanError{})
	})

	t.Run("create mutant dna with repo error", func(t *testing.T) {
		repo := new(mocks.Repository)
		ctx := context.Background()
		mutantDnaSequence := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

		hashDna := md5.Sum([]byte(strings.Join(mutantDnaSequence, "")))

		mutantDNA := domain.DNA{ID: fmt.Sprintf("%x", hashDna), DNASequence: mutantDnaSequence, Type: domain.MutantDNA}
		repo.On(saveDna, ctx, mutantDNA).Return(errors.New("error save"))

		uc := NewUseCase(repo)
		err := uc.SaveDna(ctx, mutantDnaSequence)

		assert.NoError(t, err)
	})

	t.Run("retrieve dna stats with success", func(t *testing.T) {
		repo := new(mocks.Repository)
		ctx := context.Background()
		stats := new(domain.Stats)

		repo.On(findStats, ctx).Return(stats, nil)

		uc := NewUseCase(repo)
		retrieveStats, err := uc.RetrieveStats(ctx)

		assert.Equal(t, stats, retrieveStats)
		assert.NoError(t, err)
	})

	t.Run("retrieve dna stats with fail", func(t *testing.T) {
		repo := new(mocks.Repository)
		ctx := context.Background()

		repo.On(findStats, ctx).Return(nil, errors.New("error repo"))

		uc := NewUseCase(repo)
		retrieveStats, err := uc.RetrieveStats(ctx)

		assert.Empty(t, retrieveStats)
		assert.Error(t, err)
	})
}
