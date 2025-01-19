package aggregation

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type comparisonBuilder struct {
	parent *Builder
}

func (b *comparisonBuilder) Eq(key string, expressions ...any) *Builder {
	e := bson.E{Key: EqOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) EqWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: EqOp, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Ne(key string, expressions ...any) *Builder {
	e := bson.E{Key: NeOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) NeWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: NeOp, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Gt(key string, expressions ...any) *Builder {
	e := bson.E{Key: GtOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) GtWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: GtOp, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Gte(key string, expressions ...any) *Builder {
	e := bson.E{Key: GteOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) GteWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: GteOp, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Lt(key string, expressions ...any) *Builder {
	e := bson.E{Key: LtOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) LtWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: LtOp, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Lte(key string, expressions ...any) *Builder {
	e := bson.E{Key: LteOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) LteWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: LteOp, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) IndexOfArray(key, array, filed string) *Builder {
	e := bson.E{Key: IndexOfArrayOp, Value: []any{array, filed}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}
