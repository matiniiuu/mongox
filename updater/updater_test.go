package updater

import (
	"context"
	"testing"

	mocks "github.com/matiniiuu/mongox/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/mock/gomock"
)

func TestNewUpdater(t *testing.T) {
	updater := NewUpdater[any](&mongo.Collection{})
	assert.NotNil(t, updater)
}

func TestUpdater_UpdateOne(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) IUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to update one",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().UpdateOne(ctx).Return(nil, assert.AnError).Times(1)
				return updater
			},

			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "execute successfully but modified count is 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().UpdateOne(ctx).Return(&mongo.UpdateResult{ModifiedCount: 0}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 0,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().UpdateOne(ctx).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := tc.mock(context.Background(), ctl)
			got, err := u.UpdateOne(tc.ctx)
			if tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_UpdateMany(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) IUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to update many",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().UpdateMany(ctx).Return(nil, assert.AnError).Times(1)
				return updater
			},
			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "execute successfully but modified count is 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().UpdateMany(ctx).Return(&mongo.UpdateResult{ModifiedCount: 0}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 0,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().UpdateMany(ctx).Return(&mongo.UpdateResult{ModifiedCount: 2}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 2,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := tc.mock(context.Background(), ctl)
			got, err := u.UpdateMany(tc.ctx)
			if tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_Upsert(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) IUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to upsert one",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().Upsert(ctx).Return(nil, assert.AnError).Times(1)
				return updater
			},

			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "save successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().Upsert(ctx).Return(&mongo.UpdateResult{UpsertedCount: 1}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				UpsertedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) IUpdater[any] {
				updater := mocks.NewMockIUpdater[any](ctl)
				updater.EXPECT().Upsert(ctx).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := tc.mock(context.Background(), ctl)
			got, err := u.Upsert(tc.ctx)
			if tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
