package query

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type elementQueryBuilder struct {
	parent *Builder
}

func (b *elementQueryBuilder) Exists(key string, exists bool) *Builder {
	e := bson.E{Key: ExistsOp, Value: exists}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *elementQueryBuilder) Type(key string, t bson.Type) *Builder {
	e := bson.E{Key: TypeOp, Value: t}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *elementQueryBuilder) TypeAlias(key string, alias string) *Builder {
	e := bson.E{Key: TypeOp, Value: alias}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *elementQueryBuilder) TypeArray(key string, ts ...bson.Type) *Builder {
	e := bson.E{Key: TypeOp, Value: ts}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *elementQueryBuilder) TypeArrayAlias(key string, aliases ...string) *Builder {
	e := bson.E{Key: TypeOp, Value: aliases}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}
