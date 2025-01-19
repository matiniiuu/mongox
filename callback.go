package mongox

import (
	"context"

	validator2 "github.com/go-playground/validator/v10"

	"github.com/matiniiuu/mongox/hook/validator"

	"github.com/matiniiuu/mongox/hook/model"

	"github.com/matiniiuu/mongox/callback"
	"github.com/matiniiuu/mongox/hook/field"
	"github.com/matiniiuu/mongox/operation"
)

func RegisterPlugin(name string, cb callback.CbFn, opType operation.OpType) {
	callback.Callbacks.Register(opType, name, cb)
}

func RemovePlugin(name string, opType operation.OpType) {
	callback.Callbacks.Remove(opType, name)
}

type PluginConfig struct {
	EnableDefaultFieldHook bool
	EnableModelHook        bool
	EnableValidationHook   bool
	// use to replace to the default validate instance
	Validate *validator2.Validate
}

func InitPlugin(config *PluginConfig) {
	if config.EnableDefaultFieldHook {
		opTypes := []operation.OpType{operation.OpTypeBeforeInsert, operation.OpTypeBeforeUpdate, operation.OpTypeBeforeUpsert}
		for _, opType := range opTypes {
			typ := opType
			RegisterPlugin("mongox:default_field", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
				return field.Execute(ctx, opCtx, typ, opts...)
			}, typ)
		}
	}
	if config.EnableModelHook {
		opTypes := []operation.OpType{
			operation.OpTypeBeforeInsert, operation.OpTypeAfterInsert,
			operation.OpTypeBeforeDelete, operation.OpTypeAfterDelete,
			operation.OpTypeBeforeUpdate, operation.OpTypeAfterUpdate,
			operation.OpTypeBeforeUpsert, operation.OpTypeAfterUpsert,
			operation.OpTypeBeforeFind, operation.OpTypeAfterFind,
		}
		for _, opType := range opTypes {
			typ := opType
			RegisterPlugin("mongox:model", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
				return model.Execute(ctx, opCtx, typ, opts...)
			}, typ)
		}
	}
	if config.EnableValidationHook {
		validator.SetValidate(config.Validate)
		opTypes := []operation.OpType{operation.OpTypeBeforeInsert, operation.OpTypeBeforeUpsert}
		for _, opType := range opTypes {
			typ := opType
			RegisterPlugin("mongox:validation", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
				return validator.Execute(ctx, opCtx, typ, opts...)
			}, typ)
		}
	}
}
