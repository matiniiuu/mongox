package deleter

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

//go:generate optioner -type OpContext
type OpContext struct {
	Col          *mongo.Collection `opt:"-"`
	Filter       any               `opt:"-"`
	MongoOptions any
	ModelHook    any
}

type (
	beforeHookFn func(ctx context.Context, opContext *OpContext, opts ...any) error
	afterHookFn  func(ctx context.Context, opContext *OpContext, opts ...any) error
)
