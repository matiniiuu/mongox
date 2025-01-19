package validator

import (
	"context"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/matiniiuu/mongox/operation"
)

var validate = validator.New()

func SetValidate(v *validator.Validate) {
	if v != nil {
		validate = v
	}
}

func getPayload(opCtx *operation.OpContext, opType operation.OpType) any {
	if opCtx == nil {
		return nil
	}
	switch opType {
	case operation.OpTypeBeforeInsert:
		return opCtx.Doc
	case operation.OpTypeBeforeUpsert:
		return opCtx.Replacement
	default:
		return nil
	}
}

func Execute(ctx context.Context, opCtx *operation.OpContext, opType operation.OpType, opts ...any) error {
	payLoad := getPayload(opCtx, opType)
	if payLoad == nil {
		return nil
	}
	value := reflect.ValueOf(payLoad)

	switch value.Type().Kind() {
	case reflect.Slice:
		return executeSlice(ctx, value, opts...)
	case reflect.Ptr:
		if value.IsZero() {
			return nil
		}
		return execute(ctx, value, opts...)
	default:
		return nil
	}
}

func executeSlice(ctx context.Context, docs reflect.Value, opts ...any) error {
	for i := 0; i < docs.Len(); i++ {
		doc := docs.Index(i)
		if err := execute(ctx, doc, opts...); err != nil {
			return err
		}
	}
	return nil
}

func execute(ctx context.Context, value reflect.Value, _ ...any) error {
	doc := validateStruct(value)
	if doc == nil {
		return nil
	}
	return validate.StructCtx(ctx, doc)
}

func validateStruct(doc reflect.Value) any {
	if doc.Kind() == reflect.Pointer && !doc.IsNil() {
		doc = doc.Elem()
	}
	if doc.Kind() != reflect.Struct || doc.Type().ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return nil
	}
	return doc.Interface()
}
