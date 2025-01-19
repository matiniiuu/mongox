package aggregation

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewBuilder() *Builder {
	b := &Builder{d: bson.D{}}

	b.arithmeticBuilder = arithmeticBuilder{parent: b}
	b.comparisonBuilder = comparisonBuilder{parent: b}
	b.logicalBuilder = logicalBuilder{parent: b}
	b.stringBuilder = stringBuilder{parent: b}
	b.arrayBuilder = arrayBuilder{parent: b}
	b.dateBuilder = dateBuilder{parent: b}
	b.condBuilder = condBuilder{parent: b}
	b.accumulatorsBuilder = accumulatorsBuilder{parent: b}

	return b
}

type Builder struct {
	arithmeticBuilder
	comparisonBuilder
	logicalBuilder
	stringBuilder
	arrayBuilder
	dateBuilder
	condBuilder
	accumulatorsBuilder

	d bson.D
}

func (b *Builder) Build() bson.D {
	return b.d
}

func (b *Builder) KeyValue(key string, value any) *Builder {
	b.d = append(b.d, bson.E{Key: key, Value: value})
	return b
}

// tryMergeValue attempts to merge the provided bson.E elements into an existing bson.D element
// in the builder's data slice, identified by the specified key.
func (b *Builder) tryMergeValue(key string, e ...bson.E) bool {
	for idx, datum := range b.d {
		if datum.Key == key {
			if m, ok := datum.Value.(bson.D); ok {
				m = append(m, e...)
				b.d[idx].Value = m
				return true
			}
		}
	}
	return false
}
