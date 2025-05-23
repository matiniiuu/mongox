package finder

import (
	"context"

	"github.com/matiniiuu/mongox/callback"
	"github.com/matiniiuu/mongox/operation"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
type IFinder[T any] interface {
	FindOne(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*T, error)
	Find(ctx context.Context, opts ...options.Lister[options.FindOptions]) ([]*T, error)
	Count(ctx context.Context, opts ...options.Lister[options.CountOptions]) (int64, error)
}

func NewFinder[T any](collection *mongo.Collection) *Finder[T] {
	return &Finder[T]{collection: collection, filter: bson.D{}}
}

var _ IFinder[any] = (*Finder[any])(nil)

type Finder[T any] struct {
	collection  *mongo.Collection
	filter      any
	updates     any
	modelHook   any
	beforeHooks []beforeHookFn
	afterHooks  []afterHookFn[T]
}

func (f *Finder[T]) RegisterBeforeHooks(hooks ...beforeHookFn) *Finder[T] {
	f.beforeHooks = append(f.beforeHooks, hooks...)
	return f
}

// RegisterAfterHooks is used to set the after hooks of the query
// If you register the hook for FindOne, the opContext.Docs will be nil
// If you register the hook for Find, the opContext.Doc will be nil
func (f *Finder[T]) RegisterAfterHooks(hooks ...afterHookFn[T]) *Finder[T] {
	f.afterHooks = append(f.afterHooks, hooks...)
	return f
}

// Filter is used to set the filter of the query
func (f *Finder[T]) Filter(filter any) *Finder[T] {
	f.filter = filter
	return f
}

func (f *Finder[T]) Updates(update any) *Finder[T] {
	f.updates = update
	return f
}

func (f *Finder[T]) ModelHook(modelHook any) *Finder[T] {
	f.modelHook = modelHook
	return f
}

func (f *Finder[T]) preActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opTypes ...operation.OpType) (err error) {
	for _, opType := range opTypes {
		err = callback.GetCallback().Execute(ctx, globalOpContext, opType)
		if err != nil {
			return
		}
	}
	for _, beforeHook := range f.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return
}

func (f *Finder[T]) postActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *AfterOpContext[T], opTypes ...operation.OpType) (err error) {
	for _, opType := range opTypes {
		err = callback.GetCallback().Execute(ctx, globalOpContext, opType)
		if err != nil {
			return
		}
	}
	for _, afterHook := range f.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return
		}
	}
	return
}

func (f *Finder[T]) FindOne(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*T, error) {
	t := new(T)

	globalOpContext := operation.NewOpContext(f.collection, operation.WithDoc(t), operation.WithFilter(f.filter), operation.WithMongoOptions(opts), operation.WithModelHook(f.modelHook))
	err := f.preActionHandler(ctx, globalOpContext, NewOpContext(f.collection, f.filter, WithMongoOptions(opts), WithModelHook(f.modelHook)), operation.OpTypeBeforeFind)
	if err != nil {
		return nil, err
	}

	err = f.collection.FindOne(ctx, f.filter, opts...).Decode(t)
	if err != nil {
		return nil, err
	}

	err = f.postActionHandler(ctx, globalOpContext, NewAfterOpContext[T](NewOpContext(f.collection, f.filter, WithMongoOptions(opts), WithModelHook(f.modelHook)), WithDoc(t)), operation.OpTypeAfterFind)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *Finder[T]) Find(ctx context.Context, opts ...options.Lister[options.FindOptions]) ([]*T, error) {
	t := make([]*T, 0)

	opContext := operation.NewOpContext(f.collection, operation.WithFilter(f.filter), operation.WithMongoOptions(opts), operation.WithModelHook(f.modelHook))
	err := f.preActionHandler(ctx, opContext, NewOpContext(f.collection, f.filter, WithMongoOptions(opts), WithModelHook(f.modelHook)), operation.OpTypeBeforeFind)
	if err != nil {
		return nil, err
	}

	cursor, err := f.collection.Find(ctx, f.filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &t)
	if err != nil {
		return nil, err
	}

	opContext.Doc = t
	err = f.postActionHandler(ctx, opContext, NewAfterOpContext[T](NewOpContext(f.collection, f.filter, WithMongoOptions(opts), WithModelHook(f.modelHook)), WithDocs(t)), operation.OpTypeAfterFind)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *Finder[T]) Count(ctx context.Context, opts ...options.Lister[options.CountOptions]) (int64, error) {
	return f.collection.CountDocuments(ctx, f.filter, opts...)
}

func (f *Finder[T]) Distinct(ctx context.Context, fieldName string, opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult {
	return f.collection.Distinct(ctx, fieldName, f.filter, opts...)
}

// DistinctWithParse is used to parse the result of Distinct
// result must be a pointer
func (f *Finder[T]) DistinctWithParse(ctx context.Context, fieldName string, result any, opts ...options.Lister[options.DistinctOptions]) error {
	distinctResult := f.collection.Distinct(ctx, fieldName, f.filter, opts...)
	if distinctResult.Err() != nil {
		return distinctResult.Err()
	}
	err := distinctResult.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (f *Finder[T]) FindOneAndUpdate(ctx context.Context, opts ...options.Lister[options.FindOneAndUpdateOptions]) (*T, error) {
	t := new(T)

	globalOpContext := operation.NewOpContext(f.collection, operation.WithDoc(t), operation.WithFilter(f.filter), operation.WithUpdates(f.updates), operation.WithMongoOptions(opts), operation.WithModelHook(f.modelHook))
	err := f.preActionHandler(ctx, globalOpContext, NewOpContext(f.collection, f.filter, WithUpdates(f.updates), WithMongoOptions(opts), WithModelHook(f.modelHook)), operation.OpTypeBeforeFind, operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	err = f.collection.FindOneAndUpdate(ctx, f.filter, f.updates, opts...).Decode(t)
	if err != nil {
		return nil, err
	}

	err = f.postActionHandler(ctx, globalOpContext, NewAfterOpContext[T](NewOpContext(f.collection, f.filter, WithUpdates(f.updates), WithMongoOptions(opts), WithModelHook(f.modelHook)), WithDoc(t)), operation.OpTypeAfterFind, operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}

	return t, nil
}
