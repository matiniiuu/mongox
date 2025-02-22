package field

import (
	"context"
	"reflect"

	"github.com/matiniiuu/mongox/operation"
)

func Execute(ctx context.Context, opCtx *operation.OpContext, opType operation.OpType, opts ...any) error {
	doc := opCtx.Doc
	if doc == nil {
		return nil
	}
	valueOf := reflect.ValueOf(doc)
	opts = append([]any{opCtx.Updates, opCtx.Replacement, opCtx.MongoOptions}, opts...)
	switch valueOf.Type().Kind() {
	case reflect.Slice:
		return executeSlice(ctx, valueOf, opType, opts...)
	case reflect.Ptr:
		if valueOf.IsZero() {
			return nil
		}
		return execute(ctx, doc, opType, opts...)
	default:
		return nil
	}
}

func executeSlice(ctx context.Context, docs reflect.Value, opType operation.OpType, opts ...any) error {
	for i := 0; i < docs.Len(); i++ {
		doc := docs.Index(i)
		if err := execute(ctx, doc.Interface(), opType, opts...); err != nil {
			return err
		}
	}
	return nil
}

func execute(_ context.Context, doc any, opType operation.OpType, opts ...any) error {
	if strategy, ok := strategies[opType]; ok {
		return strategy(doc, opts...)
	}
	return nil
}
