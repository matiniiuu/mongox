package mongox

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/matiniiuu/mongox/callback"
	"github.com/matiniiuu/mongox/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestRegisterPlugin_Insert(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before insert
	t.Run("before insert", func(t *testing.T) {
		RegisterPlugin("before insert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			user := opCtx.Doc.(*User)
			require.NotNil(t, user)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeInsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before insert", operation.OpTypeBeforeInsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after insert
	t.Run("after insert", func(t *testing.T) {
		RegisterPlugin("after insert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			users := opCtx.Doc.([]*User)
			require.NotNil(t, users)
			isCalled = true
			return nil
		}, operation.OpTypeAfterInsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterInsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after insert", operation.OpTypeAfterInsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterInsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Find(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before find
	t.Run("before find", func(t *testing.T) {
		RegisterPlugin("before find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeFind)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before find", operation.OpTypeBeforeFind)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after find
	t.Run("after find", func(t *testing.T) {
		RegisterPlugin("after find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			users := opCtx.Doc.([]*User)
			require.NotNil(t, users)
			isCalled = true
			return nil
		}, operation.OpTypeAfterFind)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterFind)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after find", operation.OpTypeAfterFind)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterFind)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Delete(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before delete
	t.Run("before delete", func(t *testing.T) {
		RegisterPlugin("before delete", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeDelete)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before delete", operation.OpTypeBeforeDelete)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after delete
	t.Run("after delete", func(t *testing.T) {
		RegisterPlugin("after delete", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeAfterDelete)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeAfterDelete)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after delete", operation.OpTypeAfterDelete)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterDelete)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Update(t *testing.T) {
	isCalled := false
	// before update
	t.Run("before update", func(t *testing.T) {
		RegisterPlugin("before update", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeUpdate)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before update", operation.OpTypeBeforeUpdate)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after update
	t.Run("after update", func(t *testing.T) {
		RegisterPlugin("after update", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeAfterUpdate)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeAfterUpdate)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after update", operation.OpTypeAfterUpdate)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeAfterUpdate)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Upsert(t *testing.T) {
	isCalled := false
	// before upsert
	t.Run("before upsert", func(t *testing.T) {
		RegisterPlugin("before upsert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Replacement)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeUpsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before upsert", operation.OpTypeBeforeUpsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after upsert
	t.Run("after upsert", func(t *testing.T) {
		RegisterPlugin("after upsert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Replacement)
			isCalled = true
			return nil
		}, operation.OpTypeAfterUpsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeAfterUpsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after upsert", operation.OpTypeAfterUpsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeAfterUpsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestPluginInit_EnableEnableDefaultFieldHook(t *testing.T) {
	t.Run("beforeInsert", func(t *testing.T) {
		model := &Base{}
		err := callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(model)),
			operation.OpTypeBeforeInsert,
		)
		require.Nil(t, err)
		require.Zero(t, model.ID)
		require.Zero(t, model.CreatedAt)

		cfg := &PluginConfig{
			EnableDefaultFieldHook: true,
		}
		InitPlugin(cfg)

		err = callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(model)),
			operation.OpTypeBeforeInsert,
		)
		require.Nil(t, err)
		require.NotZero(t, model.ID)
		require.NotZero(t, model.CreatedAt)
		RemovePlugin("mongox:default_field", operation.OpTypeBeforeInsert)
		RemovePlugin("mongox:default_field", operation.OpTypeBeforeUpdate)
		RemovePlugin("mongox:default_field", operation.OpTypeBeforeUpsert)
	})
	t.Run("beforeUpdate", func(t *testing.T) {
		var (
			model = &Base{}
			m     = bson.M{}
		)
		err := callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(model), operation.WithUpdates(m)),
			operation.OpTypeBeforeUpdate,
		)
		require.Nil(t, err)
		require.Zero(t, model.ID)
		require.Zero(t, model.CreatedAt)
		require.Zero(t, model.UpdatedAt)

		cfg := &PluginConfig{
			EnableDefaultFieldHook: true,
		}
		InitPlugin(cfg)

		err = callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(model), operation.WithUpdates(m)),
			operation.OpTypeBeforeUpdate,
		)
		require.Nil(t, err)
		require.Zero(t, model.ID)
		require.Zero(t, model.CreatedAt)
		require.NotZero(t, model.UpdatedAt)
		require.Equal(t, bson.M{
			"$set": bson.M{
				"updated_at": model.UpdatedAt,
			},
		}, m)

		RemovePlugin("mongox:default_field", operation.OpTypeBeforeInsert)
		RemovePlugin("mongox:default_field", operation.OpTypeBeforeUpdate)
		RemovePlugin("mongox:default_field", operation.OpTypeBeforeUpsert)
	})
	t.Run("beforeUpsert", func(t *testing.T) {
		var (
			model = &Base{}
			m     = bson.M{}
		)
		err := callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(model), operation.WithUpdates(m)),
			operation.OpTypeBeforeUpsert,
		)
		require.Nil(t, err)
		require.Zero(t, model.ID)
		require.Zero(t, model.CreatedAt)
		require.Zero(t, model.UpdatedAt)

		cfg := &PluginConfig{
			EnableDefaultFieldHook: true,
		}
		InitPlugin(cfg)

		err = callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(model), operation.WithUpdates(m)),
			operation.OpTypeBeforeUpsert,
		)
		require.Nil(t, err)
		require.NotZero(t, model.ID)
		require.NotZero(t, model.CreatedAt)
		require.NotZero(t, model.UpdatedAt)
		require.Equal(t, bson.M{
			"$set": bson.M{
				"updated_at": model.UpdatedAt,
			},
			"$setOnInsert": bson.M{
				"_id":        model.ID,
				"created_at": model.CreatedAt,
			},
		}, m)

		RemovePlugin("mongox:default_field", operation.OpTypeBeforeInsert)
		RemovePlugin("mongox:default_field", operation.OpTypeBeforeUpdate)
		RemovePlugin("mongox:default_field", operation.OpTypeBeforeUpsert)
	})
}

type testModelHookStruct int

func (t *testModelHookStruct) BeforeInsert(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) AfterInsert(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) BeforeDelete(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) AfterDelete(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) BeforeUpdate(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) AfterUpdate(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) BeforeUpsert(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) AfterUpsert(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) BeforeFind(_ context.Context) error {
	*t++
	return nil
}

func (t *testModelHookStruct) AfterFind(_ context.Context) error {
	*t++
	return nil
}

func TestPluginInit_EnableModelHook(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		ocOption func(tm *testModelHookStruct) []operation.OpContextOption
		opType   operation.OpType

		wantErr error
		want    testModelHookStruct
	}{
		{
			name: "beforeInsert",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithDoc(tm),
				}
			},
			opType:  operation.OpTypeBeforeInsert,
			wantErr: nil,
			want:    1,
		},
		{
			name: "beforeInsert with model hook",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				*tm = 1
				return []operation.OpContextOption{
					operation.WithDoc(new(testModelHookStruct)),
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeBeforeInsert,
			wantErr: nil,
			want:    2,
		},
		{
			name: "afterInsert",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithDoc(tm),
				}
			},
			opType:  operation.OpTypeAfterInsert,
			wantErr: nil,
			want:    1,
		},
		{
			name: "afterInsert with model hook",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				*tm = 1
				return []operation.OpContextOption{
					operation.WithDoc(new(testModelHookStruct)),
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeAfterInsert,
			wantErr: nil,
			want:    2,
		},
		{
			name: "beforeDelete",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeBeforeDelete,
			wantErr: nil,
			want:    1,
		},
		{
			name: "afterDelete",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeAfterDelete,
			wantErr: nil,
			want:    1,
		},
		{
			name: "beforeUpdate",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithUpdates(tm),
				}
			},
			opType:  operation.OpTypeBeforeUpdate,
			wantErr: nil,
			want:    1,
		},
		{
			name: "beforeUpdate with model hook",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				*tm = 1
				return []operation.OpContextOption{
					operation.WithUpdates(new(testModelHookStruct)),
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeBeforeUpdate,
			wantErr: nil,
			want:    2,
		},
		{
			name: "afterUpdate",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithUpdates(tm),
				}
			},
			opType:  operation.OpTypeAfterUpdate,
			wantErr: nil,
			want:    1,
		},
		{
			name: "afterUpdate with model hook",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				*tm = 1
				return []operation.OpContextOption{
					operation.WithUpdates(new(testModelHookStruct)),
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeAfterUpdate,
			wantErr: nil,
			want:    2,
		},
		{
			name: "beforeUpsert",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithUpdates(tm),
				}
			},
			opType:  operation.OpTypeBeforeUpsert,
			wantErr: nil,
			want:    1,
		},
		{
			name: "beforeUpsert with model hook",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				*tm = 1
				return []operation.OpContextOption{
					operation.WithUpdates(new(testModelHookStruct)),
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeBeforeUpsert,
			wantErr: nil,
			want:    2,
		},
		{
			name: "afterUpsert",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithUpdates(tm),
				}
			},
			opType:  operation.OpTypeAfterUpsert,
			wantErr: nil,
			want:    1,
		},
		{
			name: "afterUpsert with model hook",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				*tm = 1
				return []operation.OpContextOption{
					operation.WithUpdates(new(testModelHookStruct)),
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeAfterUpsert,
			wantErr: nil,
			want:    2,
		},
		{
			name: "beforeFind",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeBeforeFind,
			wantErr: nil,
			want:    1,
		},
		{
			name: "afterFind",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				return []operation.OpContextOption{
					operation.WithDoc(tm),
				}
			},
			opType:  operation.OpTypeAfterFind,
			wantErr: nil,
			want:    1,
		},
		{
			name: "afterFind with model hook",
			ctx:  context.Background(),
			ocOption: func(tm *testModelHookStruct) []operation.OpContextOption {
				*tm = 1
				return []operation.OpContextOption{
					operation.WithDoc(new(testModelHookStruct)),
					operation.WithModelHook(tm),
				}
			},
			opType:  operation.OpTypeAfterFind,
			wantErr: nil,
			want:    2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tm := new(testModelHookStruct)
			err := callback.GetCallback().Execute(
				tc.ctx,
				operation.NewOpContext(nil, tc.ocOption(tm)...),
				tc.opType,
			)
			require.Nil(t, err)
			cfg := &PluginConfig{
				EnableModelHook: true,
			}
			InitPlugin(cfg)
			err = callback.GetCallback().Execute(
				tc.ctx,
				operation.NewOpContext(nil, tc.ocOption(tm)...),
				tc.opType,
			)
			require.Equal(t, tc.wantErr, err)
			require.Equal(t, tc.want, *tm)
			remoteModelPlugin()
		})
	}
}

func remoteModelPlugin() {
	RemovePlugin("mongox:model", operation.OpTypeBeforeInsert)
	RemovePlugin("mongox:model", operation.OpTypeAfterInsert)
	RemovePlugin("mongox:model", operation.OpTypeBeforeDelete)
	RemovePlugin("mongox:model", operation.OpTypeAfterDelete)
	RemovePlugin("mongox:model", operation.OpTypeBeforeUpdate)
	RemovePlugin("mongox:model", operation.OpTypeAfterUpdate)
	RemovePlugin("mongox:model", operation.OpTypeBeforeUpsert)
	RemovePlugin("mongox:model", operation.OpTypeAfterUpsert)
	RemovePlugin("mongox:model", operation.OpTypeBeforeFind)
	RemovePlugin("mongox:model", operation.OpTypeAfterFind)
}

func TestPluginInit_EnableValidationHook(t *testing.T) {
	type TestModel struct {
		Name string `validate:"required"`
	}
	t.Run("beforeInsert", func(t *testing.T) {
		err := callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(&TestModel{})),
			operation.OpTypeBeforeInsert,
		)
		require.Nil(t, err)

		cfg := &PluginConfig{
			EnableValidationHook: true,
		}
		InitPlugin(cfg)

		err = callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(&TestModel{})),
			operation.OpTypeBeforeInsert,
		)
		require.NotNil(t, err.(validator.ValidationErrors))
		RemovePlugin("mongox:validation", operation.OpTypeBeforeInsert)
		RemovePlugin("mongox:validation", operation.OpTypeBeforeUpsert)
	})
	t.Run("beforeUpsert", func(t *testing.T) {
		err := callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithDoc(&TestModel{})),
			operation.OpTypeBeforeUpsert,
		)
		require.Nil(t, err)

		cfg := &PluginConfig{
			EnableValidationHook: true,
		}
		InitPlugin(cfg)

		err = callback.GetCallback().Execute(
			context.Background(),
			operation.NewOpContext(nil, operation.WithReplacement(&TestModel{})),
			operation.OpTypeBeforeUpsert,
		)
		require.NotNil(t, err.(validator.ValidationErrors))
		RemovePlugin("mongox:validation", operation.OpTypeBeforeInsert)
		RemovePlugin("mongox:validation", operation.OpTypeBeforeUpsert)
	})
}
