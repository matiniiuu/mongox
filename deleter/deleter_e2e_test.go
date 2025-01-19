//go:build e2e

package deleter

import (
	"context"
	"fmt"
	"testing"

	"github.com/matiniiuu/mongox/internal/pkg/utils"

	"github.com/matiniiuu/mongox/callback"
	"github.com/matiniiuu/mongox/operation"

	"github.com/stretchr/testify/require"

	"github.com/matiniiuu/mongox/bsonx"

	"github.com/matiniiuu/mongox/builder/query"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type testTempUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

func newCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	require.NoError(t, err)
	require.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := client.Database("db-test").Collection("test_user")
	return collection
}

func TestDeleter_e2e_New(t *testing.T) {
	result := NewDeleter[any](newCollection(t))
	require.NotNil(t, result, "Expected non-nil Deleter")
}

func TestDeleter_e2e_DeleteOne(t *testing.T) {
	collection := newCollection(t)
	deleter := NewDeleter[testTempUser](collection)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter     bson.D
		opts       []options.Lister[options.DeleteOneOptions]
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn

		ctx       context.Context
		want      *mongo.DeleteResult
		wantError require.ErrorAssertionFunc
	}{
		{
			name:      "error: nil filter",
			before:    func(_ context.Context, _ *testing.T) {},
			after:     func(_ context.Context, _ *testing.T) {},
			filter:    nil,
			ctx:       context.Background(),
			opts:      []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			want:      nil,
			wantError: require.Error,
		},
		{
			name: "deleted count: 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			filter: query.NewBuilder().Id("2").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name: "delete success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("before hook error")
					},
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name:   "global after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("after hook error")
					},
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return fmt.Errorf("before hook error")
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name:   "after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return fmt.Errorf("after hook error")
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).DeleteOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			deleter.beforeHooks = nil
			deleter.afterHooks = nil
		})
	}
}

func TestDeleter_e2e_DeleteMany(t *testing.T) {
	collection := newCollection(t)
	deleter := NewDeleter[testTempUser](collection)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter     any
		opts       []options.Lister[options.DeleteManyOptions]
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn

		ctx       context.Context
		want      *mongo.DeleteResult
		wantError require.ErrorAssertionFunc
	}{
		{
			name:      "error: nil filter",
			before:    func(_ context.Context, _ *testing.T) {},
			after:     func(_ context.Context, _ *testing.T) {},
			filter:    nil,
			ctx:       context.Background(),
			opts:      []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want:      nil,
			wantError: require.Error,
		},
		{
			name: "deleted count: 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name: "delete success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: bsonx.M("name", "Mingyong Chen"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("before hook error")
					},
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("after hook error")
					},
				},
			},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want:   nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return fmt.Errorf("before hook error")
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return fmt.Errorf("after hook error")
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want:   nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).DeleteMany(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			deleter.beforeHooks = nil
			deleter.afterHooks = nil
		})
	}
}
