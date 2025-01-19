package creator

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

//go:generate optioner -type OpContext
type OpContext[T any] struct {
	Col          *mongo.Collection `opt:"-"`
	Doc          *T
	Docs         []*T
	MongoOptions any
	ModelHook    any
}

type (
	hookFn[T any] func(ctx context.Context, opContext *OpContext[T], opts ...any) error
)
