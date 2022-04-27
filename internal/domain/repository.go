package domain

import "context"

type Repository interface {
	SaveDna(ctx context.Context, mutant Mutant) error
	FindStats(ctx context.Context) (*Stats, error)
}
