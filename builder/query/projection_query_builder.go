package query

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type projectionQueryBuilder struct {
	parent *Builder
}

func (b *projectionQueryBuilder) Slice(key string, number int) *Builder {
	e := bson.E{Key: SliceOp, Value: number}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *projectionQueryBuilder) SliceRanger(key string, start, end int) *Builder {
	e := bson.E{Key: SliceOp, Value: []int{start, end}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}
