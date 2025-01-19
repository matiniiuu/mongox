package aggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestBuilder_KeyValue(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "name", Value: "chenmingyong"}}, NewBuilder().KeyValue("name", "chenmingyong").Build())
}

func TestBuilder_tryMergeValue(t *testing.T) {
	assert.True(t, NewBuilder().Push("items", "$item").tryMergeValue("items", bson.E{Key: "$avg", Value: "$items"}))
}
