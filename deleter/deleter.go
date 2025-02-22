package deleter

import (
	"context"

	"github.com/matiniiuu/mongox/callback"
	"github.com/matiniiuu/mongox/operation"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=deleter.go -destination=../mock/deleter.mock.go -package=mocks
type IDeleter[T any] interface {
	DeleteOne(ctx context.Context, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error)
}

var _ IDeleter[any] = (*Deleter[any])(nil)

func NewDeleter[T any](collection *mongo.Collection) *Deleter[T] {
	return &Deleter[T]{collection: collection, filter: nil}
}

type Deleter[T any] struct {
	collection  *mongo.Collection
	filter      any
	modelHook   any
	beforeHooks []beforeHookFn
	afterHooks  []afterHookFn
}

func (d *Deleter[T]) RegisterBeforeHooks(hooks ...beforeHookFn) *Deleter[T] {
	d.beforeHooks = append(d.beforeHooks, hooks...)
	return d
}

func (d *Deleter[T]) RegisterAfterHooks(hooks ...afterHookFn) *Deleter[T] {
	d.afterHooks = append(d.afterHooks, hooks...)
	return d
}

func (d *Deleter[T]) preActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := callback.GetCallback().Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, beforeHook := range d.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Deleter[T]) postActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := callback.GetCallback().Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range d.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

// Filter is used to set the filter of the query
func (d *Deleter[T]) Filter(filter any) *Deleter[T] {
	d.filter = filter
	return d
}

func (d *Deleter[T]) ModelHook(modelHook any) *Deleter[T] {
	d.modelHook = modelHook
	return d
}

func (d *Deleter[T]) DeleteOne(ctx context.Context, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	globalPoContext := operation.NewOpContext(d.collection, operation.WithFilter(d.filter), operation.WithMongoOptions(opts), operation.WithModelHook(d.modelHook))
	err := d.preActionHandler(ctx, globalPoContext, NewOpContext(d.collection, d.filter, WithMongoOptions(opts), WithModelHook(d.modelHook)), operation.OpTypeBeforeDelete)
	if err != nil {
		return nil, err
	}

	result, err := d.collection.DeleteOne(ctx, d.filter, opts...)
	if err != nil {
		return nil, err
	}

	err = d.postActionHandler(ctx, globalPoContext, NewOpContext(d.collection, d.filter, WithMongoOptions(opts), WithModelHook(d.modelHook)), operation.OpTypeAfterDelete)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Deleter[T]) DeleteMany(ctx context.Context, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error) {
	globalPoContext := operation.NewOpContext(d.collection, operation.WithFilter(d.filter), operation.WithMongoOptions(opts), operation.WithModelHook(d.modelHook))
	err := d.preActionHandler(ctx, globalPoContext, NewOpContext(d.collection, d.filter, WithMongoOptions(opts), WithModelHook(d.modelHook)), operation.OpTypeBeforeDelete)
	if err != nil {
		return nil, err
	}

	result, err := d.collection.DeleteMany(ctx, d.filter, opts...)
	if err != nil {
		return nil, err
	}

	err = d.postActionHandler(ctx, globalPoContext, NewOpContext(d.collection, d.filter, WithMongoOptions(opts), WithModelHook(d.modelHook)), operation.OpTypeAfterDelete)
	if err != nil {
		return nil, err
	}

	return result, nil
}
