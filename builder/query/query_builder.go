package query

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewBuilder() *Builder {
	query := &Builder{
		data: bson.D{},
		err:  make([]error, 0),
	}
	query.comparisonQueryBuilder = comparisonQueryBuilder{parent: query}
	query.logicalQueryBuilder = logicalQueryBuilder{parent: query}
	query.elementQueryBuilder = elementQueryBuilder{parent: query}
	query.arrayQueryBuilder = arrayQueryBuilder{parent: query}
	query.evaluationQueryBuilder = evaluationQueryBuilder{parent: query}
	query.projectionQueryBuilder = projectionQueryBuilder{parent: query}
	return query
}

type Builder struct {
	data bson.D
	comparisonQueryBuilder
	logicalQueryBuilder
	elementQueryBuilder
	arrayQueryBuilder
	evaluationQueryBuilder
	projectionQueryBuilder

	err []error
}

func (b *Builder) Build() bson.D {
	return b.data
}

// Id appends an element with '_id' key and given value to the builder's data slice.
func (b *Builder) Id(v any) *Builder {
	b.data = append(b.data, bson.E{Key: IdOp, Value: v})
	return b
}

// KeyValue appends given key-value pair to the builder's data slice.
func (b *Builder) KeyValue(key string, value any) *Builder {
	b.data = append(b.data, bson.E{Key: key, Value: value})
	return b
}

// tryMergeValue attempts to merge the provided bson.E elements into an existing bson.D element
// in the builder's data slice, identified by the specified key.
func (b *Builder) tryMergeValue(key string, e ...bson.E) bool {
	for idx, datum := range b.data {
		if datum.Key == key {
			if m, ok := datum.Value.(bson.D); ok {
				m = append(m, e...)
				b.data[idx].Value = m
				return true
			}
		}
	}
	return false
}
