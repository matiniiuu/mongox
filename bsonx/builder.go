package bsonx

import "go.mongodb.org/mongo-driver/v2/bson"

// DBuilder is a builder for bson.D
type DBuilder struct {
	d bson.D
}

func NewD() *DBuilder {
	return &DBuilder{d: bson.D{}}
}

func (b *DBuilder) Add(key string, value any) *DBuilder {
	b.d = append(b.d, bson.E{Key: key, Value: value})
	return b
}

func (b *DBuilder) Build() bson.D {
	return b.d
}
