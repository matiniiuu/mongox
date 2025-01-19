package update

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestBuilder_KeyValue(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 18}, bson.E{Key: "scores", Value: []int{100, 99, 98}}}, NewBuilder().KeyValue("name", "cmy").KeyValue("age", 18).KeyValue("scores", []int{100, 99, 98}).Build())
}
