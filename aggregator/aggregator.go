package aggregator

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=aggregator.go -destination=../mock/aggregator.mock.go -package=mocks
type IAggregator[T any] interface {
	Aggregate(ctx context.Context, opts ...options.Lister[options.AggregateOptions]) ([]*T, error)
	AggregateWithParse(ctx context.Context, result any, opts ...options.Lister[options.AggregateOptions]) error
}

var _ IAggregator[any] = (*Aggregator[any])(nil)

type Aggregator[T any] struct {
	collection *mongo.Collection
	pipeline   any
}

func NewAggregator[T any](collection *mongo.Collection) *Aggregator[T] {
	return &Aggregator[T]{
		collection: collection,
	}
}

func (a *Aggregator[T]) Pipeline(pipeline any) *Aggregator[T] {
	a.pipeline = pipeline
	return a
}

func (a *Aggregator[T]) Aggregate(ctx context.Context, opts ...options.Lister[options.AggregateOptions]) ([]*T, error) {
	cursor, err := a.collection.Aggregate(ctx, a.pipeline, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]*T, 0)
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AggregateWithParse is used to parse the result of the aggregation
// result must be a pointer to a slice
func (a *Aggregator[T]) AggregateWithParse(ctx context.Context, result any, opts ...options.Lister[options.AggregateOptions]) error {
	cursor, err := a.collection.Aggregate(ctx, a.pipeline, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, result)
	if err != nil {
		return err
	}
	return nil
}
