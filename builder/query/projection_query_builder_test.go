package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_projectionQueryBuilder_Slice(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "key", Value: bson.D{bson.E{Key: "$slice", Value: 1}}}}, NewBuilder().Slice("key", 1).Build())
}

func Test_projectionQueryBuilder_SliceRanger(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "key", Value: bson.D{bson.E{Key: "$slice", Value: []int{1, 2}}}}}, NewBuilder().SliceRanger("key", 1, 2).Build())
}
