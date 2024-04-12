package seed

import (
	"context"

	"gorm.io/gorm"
)

// Seed is a struct that contains the data to be seeded.
type Seed struct {
	Name string
	Run  SeedFunc
}

// SeedFunc is a function that seeds data.
type SeedFunc func(*gorm.DB) error

// Seeds is a slice of Seed.
type Seeds []Seed

// All returns all the seeds.
func (s Seeds) All() []Seed {
	return s
}

// Seeder is an interface for seeding data.
type Seeder interface {
	Seed(ctx context.Context, seeds ...Seed) error
}

type seederImpl struct {
	db *gorm.DB
}

// Opt is a function that configures the Seeder.
type Opt func(*seederImpl)

// WithDatabase sets the database for the Seeder.
func WithDatabase(db *gorm.DB) Opt {
	return func(s *seederImpl) {
		s.db = db
	}
}

// NewSeeder creates a new Seeder.
func NewSeeder(opts ...Opt) Seeder {
	s := new(seederImpl)

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Add adds a seed function to the seeder.
func (s *seederImpl) Seed(ctx context.Context, seeds ...Seed) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, seed := range seeds {
			if err := seed.Run(tx); err != nil {
				return err
			}
		}

		return nil
	})
}
