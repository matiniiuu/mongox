package finder

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

//go:generate optioner -type OpContext
type OpContext struct {
	Col          *mongo.Collection `opt:"-"`
	Filter       any               `opt:"-"`
	Updates      any
	MongoOptions any
	ModelHook    any
}

//go:generate optioner -type AfterOpContext
type AfterOpContext[T any] struct {
	*OpContext `opt:"-"`
	Doc        *T
	Docs       []*T
}

type (
	beforeHookFn       func(ctx context.Context, opContext *OpContext, opts ...any) error
	afterHookFn[T any] func(ctx context.Context, opContext *AfterOpContext[T], opts ...any) error
)

type TestUser struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Age          int64
	UnknownField string    `bson:"-"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func (tu *TestUser) DefaultCreatedAt() {
	if tu.CreatedAt.IsZero() {
		tu.CreatedAt = time.Now().Local()
	}
}

func (tu *TestUser) DefaultUpdatedAt() {
	tu.UpdatedAt = time.Now().Local()
}

type TestTempUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

type IllegalUser struct {
	ID   bson.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name"`
	Age  string
}

type UpdatedUser struct {
	Name string `bson:"name"`
	Age  int64
}

type UserName struct {
	Name string `bson:"name"`
}
