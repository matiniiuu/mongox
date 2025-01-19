package aggregation

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type stringBuilder struct {
	parent *Builder
}

func (b *stringBuilder) Concat(key string, expressions ...any) *Builder {
	e := bson.E{Key: ConcatOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ConcatWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: ConcatOp, Value: expressions})
	return b.parent
}

func (b *stringBuilder) SubstrBytes(key string, stringExpression string, byteIndex int64, byteCount int64) *Builder {
	e := bson.E{Key: SubstrBytesOp, Value: []any{stringExpression, byteIndex, byteCount}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) SubstrBytesWithoutKey(stringExpression string, byteIndex int64, byteCount int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: SubstrBytesOp, Value: []any{stringExpression, byteIndex, byteCount}})
	return b.parent
}

func (b *stringBuilder) ToLower(key string, expression any) *Builder {
	e := bson.E{Key: ToLowerOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ToLowerWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: ToLowerOp, Value: expression})
	return b.parent
}

func (b *stringBuilder) ToUpper(key string, expression any) *Builder {
	e := bson.E{Key: ToUpperOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ToUpperWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: ToUpperOp, Value: expression})
	return b.parent
}

func (b *stringBuilder) Contact(key string, expressions ...any) *Builder {
	e := bson.E{Key: ContactOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ContactWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: ContactOp, Value: expressions})
	return b.parent
}
