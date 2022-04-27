package domain

import "context"

type Repository interface {
	SaveDna(ctx context.Context, dna DNA) error
	FindStats(ctx context.Context) (*Stats, error)
}
