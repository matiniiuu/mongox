package update

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewBuilder() *Builder {
	b := &Builder{data: bson.D{}}
	b.fieldUpdateBuilder = fieldUpdateBuilder{parent: b}
	b.arrayUpdateBuilder = arrayUpdateBuilder{parent: b}
	return b
}

type Builder struct {
	data bson.D
	fieldUpdateBuilder
	arrayUpdateBuilder
}

// KeyValue appends given key-value pair to the builder's data slice.
func (b *Builder) KeyValue(key string, value any) *Builder {
	b.data = append(b.data, bson.E{Key: key, Value: value})
	return b
}

func (b *Builder) Build() bson.D {
	return b.data
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
