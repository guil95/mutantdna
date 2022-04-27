package repositories

import (
	"context"
	"log"
	"math"
	"strings"

	"github.com/guil95/mutantdna/internal/domain"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) SaveDna(ctx context.Context, dna domain.DNA) error {
	_, err := r.db.Exec(
		"Insert into dna (id, dna, dna_type) VALUES ($1, $2, $3)",
		dna.ID,
		strings.Join(dna.DNASequence, ""),
		dna.Type,
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
	}

	err := r.db.Get(&dest, `
	SELECT
		COUNT(CASE WHEN dna_type = 'Mutant' THEN dna.id END) AS mutant_type,
		COUNT(CASE WHEN dna_type = 'Human' THEN dna.id END) AS human_type
	FROM dna`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ratio := float64(dest.MutantType)

	if dest.HumanType > 0 {
		ratio = math.Round((float64(dest.MutantType) / float64(dest.HumanType))*100) / 100
	}

	return &domain.Stats{
		CountHumanDNA:  dest.HumanType,
		CountMutantDNA: dest.MutantType,
		Ratio:          ratio,
	}, nil
}
