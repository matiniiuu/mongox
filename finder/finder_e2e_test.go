//go:build e2e

package finder

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/matiniiuu/mongox/builder/update"

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

func getCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	require.NoError(t, err)
	require.NoError(t, client.Ping(context.Background(), readpref.Primary()))
	return client.Database("db-test").Collection("test_user")
}

func TestFinder_e2e_New(t *testing.T) {
	collection := getCollection(t)

	result := NewFinder[TestUser](collection)
	require.NotNil(t, result, "Expected non-nil Finder")
	require.Equal(t, collection, result.collection, "Expected finder field to be initialized correctly")
}

func TestFinder_e2e_FindOne(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[TestUser](collection)

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
		opts       []options.Lister[options.FindOneOptions]
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn[TestUser]

		ctx     context.Context
		want    *TestUser
		wantErr error
	}{
		{
			name: "no document",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "burt"),
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "find by name",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			want: &TestUser{
				Name: "Mingyong Chen",
				Age:  24,
			},
		},
		{
			name: "ignore age field",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			opts: []options.Lister[options.FindOneOptions]{
				options.FindOne().SetProjection(bsonx.M("age", 0)),
			},
			want: &TestUser{
				Name: "Mingyong Chen",
			},
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("global before hook error")
					},
				},
			},
			wantErr: errors.New("global before hook error"),
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("global after hook error")
					},
				},
			},
			wantErr: errors.New("global after hook error"),
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter.(bson.D)[0].Key != "name" || opCtx.Filter.(bson.D)[0].Value.(bson.D)[0].Value != "Mingyong Chen" {
							return errors.New("filter error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						user := opCtx.Doc.(*TestUser)
						if user.Name != "Mingyong Chen" || user.Age != 18 {
							return errors.New("result error")
						}
						return nil
					},
				},
			},
			want: &TestUser{
				Name: "Mingyong Chen",
				Age:  18,
			},
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantErr: errors.New("before hook error"),
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			afterHook: []afterHookFn[TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[TestUser], opts ...any) error {
					return errors.New("after hook error")
				},
			},
			wantErr: errors.New("after hook error"),
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter.(bson.D)[0].Key != "name" || opCtx.Filter.(bson.D)[0].Value.(bson.D)[0].Value != "Mingyong Chen" {
						return errors.New("filter error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn[TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[TestUser], opts ...any) error {
					user := opCtx.Doc
					if user.Name != "Mingyong Chen" || user.Age != 18 {
						return errors.New("after error")
					}
					return nil
				},
			},
			want: &TestUser{
				Name: "Mingyong Chen",
				Age:  18,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			user, err := finder.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).
				FindOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			require.Equal(t, tc.wantErr, err)
			if err == nil {
				tc.want.ID = user.ID
				require.Equal(t, tc.want, user)
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			finder.beforeHooks = nil
			finder.afterHooks = nil
		})
	}
}

func TestFinder_e2e_Find(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[TestUser](collection)

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
		opts       []options.Lister[options.FindOptions]
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn[TestUser]

		ctx     context.Context
		want    []*TestUser
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "nil filter error",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			wantErr: require.Error,
		},
		{
			name: "decode error",
			ctx:  context.Background(),
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&IllegalUser{
						Name: "Mingyong Chen",
						Age:  "24",
					},
					&IllegalUser{
						Name: "burt",
						Age:  "25",
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter:  bson.D{},
			wantErr: require.Error,
		},
		{
			name: "returns empty documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  24,
					},
					&TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			ctx:     context.Background(),
			filter:  query.Eq("name", "cmy"),
			want:    []*TestUser{},
			wantErr: require.NoError,
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  24,
					},
					&TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			ctx:    context.Background(),
			filter: bson.D{},
			want: []*TestUser{
				{
					Name: "Mingyong Chen",
					Age:  24,
				},
				{
					Name: "burt",
					Age:  25,
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "find by multiple name",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  24,
					},
					&TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			ctx:    context.Background(),
			filter: query.In("name", "Mingyong Chen", "burt"),
			want: []*TestUser{
				{
					Name: "Mingyong Chen",
					Age:  24,
				},
				{
					Name: "burt",
					Age:  25,
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "ignore age field",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  24,
					},
					&TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			ctx:    context.Background(),
			filter: query.In("name", "Mingyong Chen", "burt"),
			opts: []options.Lister[options.FindOptions]{
				options.Find().SetProjection(bsonx.M("age", 0)),
			},
			want: []*TestUser{
				{
					Name: "Mingyong Chen",
				},
				{
					Name: "burt",
				},
			},
			wantErr: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("before hook error")
					},
				},
			},
			ctx: context.Background(),
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			ctx:    context.Background(),
			filter: query.In("name", "Mingyong Chen", "burt"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("after hook error")
					},
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			ctx:    context.Background(),
			filter: query.In("name", "Mingyong Chen", "burt"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return errors.New("filter error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						fmt.Println(opCtx.Doc)
						users := opCtx.Doc.([]*TestUser)
						if len(users) != 2 {
							return errors.New("result error")
						}
						return nil
					},
				},
			},
			wantErr: require.NoError,
			want: []*TestUser{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return errors.New("before hook error")
				},
			},
			ctx: context.Background(),
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			ctx:    context.Background(),
			filter: query.In("name", "Mingyong Chen", "burt"),
			afterHook: []afterHookFn[TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[TestUser], opts ...any) error {
					return errors.New("after hook error")
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			ctx: context.Background(),
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: query.In("name", "Mingyong Chen", "burt"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return errors.New("filter error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn[TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[TestUser], opts ...any) error {
					users := opCtx.Docs
					if len(users) != 2 {
						return errors.New("result error")
					}
					return nil
				},
			},
			wantErr: require.NoError,
			want: []*TestUser{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			users, err := finder.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).Find(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.Equal(t, len(tc.want), len(users))
				for _, user := range users {
					var zero bson.ObjectID
					user.ID = zero
				}
				require.ElementsMatch(t, tc.want, users)
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			finder.beforeHooks = nil
			finder.afterHooks = nil
		})
	}
}

func TestFinder_e2e_Count(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter any
		opts   []options.Lister[options.CountOptions]

		ctx     context.Context
		want    int64
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "nil filter error",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			filter:  nil,
			wantErr: require.Error,
		},
		{
			name:    "returns 0",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			filter:  bson.D{},
			wantErr: require.NoError,
		},
		{
			name: "returns 1",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			opts: []options.Lister[options.CountOptions]{
				options.Count().SetComment("test"),
			},
			filter:  bson.D{},
			want:    1,
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			count, err := finder.Filter(tc.filter).Count(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.Equal(t, tc.want, count)
			}
		})
	}
}

func TestFinder_e2e_Distinct(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		fieldName string
		filter    any
		opts      []options.Lister[options.DistinctOptions]

		ctx     context.Context
		want    []string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:   "nil filter error",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			filter:  "name",
			ctx:     context.Background(),
			want:    nil,
			wantErr: require.Error,
		},
		{
			name:      "returns empty documents",
			before:    func(ctx context.Context, t *testing.T) {},
			after:     func(ctx context.Context, t *testing.T) {},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			want:      []string{},
			wantErr:   require.NoError,
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []*TestUser{
					{
						Name: "Mingyong Chen",
						Age:  24,
					},
					{
						Name: "burt",
						Age:  45,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			want: []string{
				"Mingyong Chen",
				"burt",
			},
			wantErr: require.NoError,
		},
		{
			name: "name distinct",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []*TestUser{
					{
						Name: "Mingyong Chen",
						Age:  24,
					},
					{
						Name: "Mingyong Chen",
						Age:  25,
					},
					{
						Name: "burt",
						Age:  26,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 3)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(3), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			want: []string{
				"Mingyong Chen",
				"burt",
			},
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			distinctResult := finder.Filter(tc.filter).Distinct(tc.ctx, tc.fieldName, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, distinctResult.Err())
			if distinctResult.Err() == nil {
				result := make([]string, 0)
				err := distinctResult.Decode(&result)
				require.NoError(t, err)
				require.ElementsMatch(t, tc.want, result)
			}
		})
	}
}

func TestFinder_e2e_DistinctWithParse(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		fieldName string
		filter    any
		result    []string
		opts      []options.Lister[options.DistinctOptions]

		ctx     context.Context
		want    []string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:   "nil filter error",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			filter:  "name",
			ctx:     context.Background(),
			wantErr: require.Error,
		},
		{
			name:      "returns empty documents",
			before:    func(ctx context.Context, t *testing.T) {},
			after:     func(ctx context.Context, t *testing.T) {},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			result:    []string{},
			want:      []string{},
			wantErr:   require.NoError,
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]*TestUser{
					{
						Name: "Mingyong Chen",
						Age:  24,
					},
					{
						Name: "burt",
						Age:  45,
					},
				}...))
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			result:    []string{},
			want: []string{
				"Mingyong Chen",
				"burt",
			},
			wantErr: require.NoError,
		},
		{
			name: "name distinct",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]*TestUser{
					{
						Name: "Mingyong Chen",
						Age:  24,
					},
					{
						Name: "Mingyong Chen",
						Age:  25,
					},
					{
						Name: "burt",
						Age:  26,
					},
				}...))
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 3)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(3), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			result:    []string{},
			want: []string{
				"Mingyong Chen",
				"burt",
			},
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			err := finder.Filter(tc.filter).DistinctWithParse(tc.ctx, tc.fieldName, &tc.result, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.ElementsMatch(t, tc.want, tc.result)
			}
		})
	}
}

func TestFinder_e2e_FindOneAndUpdate(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[TestUser](collection)

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
		updates    any
		opts       []options.Lister[options.FindOneAndUpdateOptions]
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn[TestUser]

		ctx     context.Context
		want    *TestUser
		wantErr error
	}{
		{
			name: "nil document",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "burt"),
			wantErr: mongo.ErrNilDocument,
		},
		{
			name: "find by name and update age",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "Mingyong Chen"),
			updates: update.Set("age", 24),
			opts:    []options.Lister[options.FindOneAndUpdateOptions]{options.FindOneAndUpdate().SetReturnDocument(options.After)},
			want: &TestUser{
				Name: "Mingyong Chen",
				Age:  24,
			},
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("global before hook error")
					},
				},
			},
			wantErr: errors.New("global before hook error"),
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "Mingyong Chen"),
			updates: update.Set("age", 24),
			opts:    []options.Lister[options.FindOneAndUpdateOptions]{options.FindOneAndUpdate().SetReturnDocument(options.After)},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("global after hook error")
					},
				},
			},
			wantErr: errors.New("global after hook error"),
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "Mingyong Chen"),
			updates: update.Set("age", 24),
			opts:    []options.Lister[options.FindOneAndUpdateOptions]{options.FindOneAndUpdate().SetReturnDocument(options.After)},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter.(bson.D)[0].Key != "name" || opCtx.Filter.(bson.D)[0].Value.(bson.D)[0].Value != "Mingyong Chen" {
							return errors.New("filter error")
						}
						if opCtx.Updates.(bson.D)[0].Value.(bson.D)[0].Key != "age" || opCtx.Updates.(bson.D)[0].Value.(bson.D)[0].Value != 24 {
							return errors.New("updates error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						user := opCtx.Doc.(*TestUser)
						if user.Name != "Mingyong Chen" || user.Age != 24 {
							return errors.New("result error")
						}
						return nil
					},
				},
			},
			want: &TestUser{
				Name: "Mingyong Chen",
				Age:  24,
			},
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantErr: errors.New("before hook error"),
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "Mingyong Chen"),
			updates: update.Set("age", 24),
			opts:    []options.Lister[options.FindOneAndUpdateOptions]{options.FindOneAndUpdate().SetReturnDocument(options.After)},
			afterHook: []afterHookFn[TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[TestUser], opts ...any) error {
					return errors.New("after hook error")
				},
			},
			wantErr: errors.New("after hook error"),
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "Mingyong Chen"),
			updates: update.Set("age", 24),
			opts:    []options.Lister[options.FindOneAndUpdateOptions]{options.FindOneAndUpdate().SetReturnDocument(options.After)},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter.(bson.D)[0].Key != "name" || opCtx.Filter.(bson.D)[0].Value.(bson.D)[0].Value != "Mingyong Chen" {
						return errors.New("filter error")
					}
					if opCtx.Updates.(bson.D)[0].Value.(bson.D)[0].Key != "age" || opCtx.Updates.(bson.D)[0].Value.(bson.D)[0].Value != 24 {
						return errors.New("updates error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn[TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[TestUser], opts ...any) error {
					user := opCtx.Doc
					if user.Name != "Mingyong Chen" || user.Age != 24 {
						return errors.New("after error")
					}
					return nil
				},
			},
			want: &TestUser{
				Name: "Mingyong Chen",
				Age:  24,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			user, err := finder.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).Updates(tc.updates).
				FindOneAndUpdate(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			require.Equal(t, tc.wantErr, err)
			if err == nil {
				tc.want.ID = user.ID
				require.Equal(t, tc.want, user)
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			finder.beforeHooks = nil
			finder.afterHooks = nil
		})
	}
}
