package aggregation

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type logicalBuilder struct {
	parent *Builder
}

func (b *logicalBuilder) And(key string, expressions ...any) *Builder {
	e := bson.E{Key: AndOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *logicalBuilder) AndWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: AndOp, Value: expressions})
	return b.parent
}

func (b *logicalBuilder) Not(key string, expressions ...any) *Builder {
	e := bson.E{Key: NotOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *logicalBuilder) NotWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: NotOp, Value: expressions})
	return b.parent
}

func (b *logicalBuilder) Or(key string, expressions ...any) *Builder {
	e := bson.E{Key: OrOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *logicalBuilder) OrWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: OrOp, Value: expressions})
	return b.parent
}
