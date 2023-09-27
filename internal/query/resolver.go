package query

import (
	"context"
	"github.com/alloykh/pack-calc/internal/models"
	"github.com/alloykh/pack-calc/pkg/log"
)

type PacksCalculator interface {
	Calculate(input int) (resp map[int]int, total int, err error)
}

type Resolver struct {
	logger    log.Logger
	packsCalc PacksCalculator
}

func NewQueryResolver(logger log.Logger, packsCalc PacksCalculator) *Resolver {
	return &Resolver{
		logger:    logger,
		packsCalc: packsCalc,
	}
}

func (r *Resolver) CalcPacks(ctx context.Context, input models.PacksRequest) (output *models.PacksResponse, err error) {

	r.logger.Infof("INPUT: %v\n", input)

	m, total, err := r.packsCalc.Calculate(input.Amount)

	if err != nil {
		r.logger.Error(err)
		return
	}

	output = &models.PacksResponse{
		Result: m,
		Total:  total,
	}

	return
}
