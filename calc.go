package calcapi

import (
	"context"
	"log"

	calc "github.com/dragonator/goa-framework-service/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res *calc.Multiplyresponse, err error) {
	s.logger.Print("calc.multiply")

	multiple := p.A * p.B

	return &calc.Multiplyresponse{
		Multiple: multiple,
	}, nil
}
