package seed

import (
	"context"
	"github.com/oktopriima/marvel/src/app/modules/base/models"
)

type SeederInterface interface {
	SecureTable
	RunSeeder
}

type SecureTable interface {
	CheckRow(ctx context.Context, m models.Model) bool
}

type RunSeeder interface {
	Run(ctx context.Context, m []models.Model) error
}
