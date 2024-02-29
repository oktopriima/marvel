package seed

import (
	"context"
	"github.com/oktopriima/marvel/app/modules/base/model"
)

type SeederInterface interface {
	SecureTable
	RunSeeder
}

type SecureTable interface {
	CheckRow(ctx context.Context, m model.Model) bool
}

type RunSeeder interface {
	Run(ctx context.Context, m []model.Model) error
}
