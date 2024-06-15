package usecases

import "context"

type SaveClimateDataUseCasePort interface {
	Execute(ctx context.Context, eventMessage []byte) (err error)
}
