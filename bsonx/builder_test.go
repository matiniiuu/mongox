package bsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestDBuilder(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "name", Value: "chenmingyong"}, bson.E{Key: "age", Value: 24}}, NewD().Add("name", "chenmingyong").Add("age", 24).Build())
}
