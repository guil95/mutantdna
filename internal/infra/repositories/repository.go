package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"math"
	"strings"

	"github.com/guil95/mutantdna/internal/domain"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) SaveDna(ctx context.Context, mutant domain.Mutant) error {
	_, err := r.db.Exec(
		"Insert into dna (id, dna, dna_type) VALUES ($1, $2, $3)",
		mutant.ID,
		strings.Join(mutant.DNA, ""),
		mutant.Type,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindStats(ctx context.Context) (*domain.Stats, error) {
	var dest struct {
		MutantType int64   `db:"mutant_type"`
		HumanType  int64   `db:"human_type"`
		Ratio      float64 `db:"ratio"`
	}

	err := r.db.Get(&dest, `
	SELECT 
		 COUNT(CASE WHEN dna_type = 'Mutant' THEN dna.id END) AS mutant_type,
		 COUNT(CASE WHEN dna_type = 'Human' THEN dna.id END) AS human_type, 
		 float4(COUNT(CASE WHEN dna_type = 'Mutant' THEN dna.id END)) / 
		 float4(COUNT(CASE WHEN dna_type = 'Human' THEN dna.id END)) 
	AS ratio FROM dna`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &domain.Stats{
		CountHumanDNA:  dest.HumanType,
		CountMutantDNA: dest.MutantType,
		Ratio:          math.Round(dest.Ratio*100) / 100,
	}, nil
}
