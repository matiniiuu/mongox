package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_logicalQueryBuilder_And(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$and", Value: []any{bson.D{bson.E{Key: "name", Value: "cmy"}}}}}, NewBuilder().And(bson.D{{Key: "name", Value: "cmy"}}).Build())
}

func Test_logicalQueryBuilder_Not(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$not", Value: bson.D{{Key: "name", Value: "cmy"}}}}, NewBuilder().Not(bson.D{{Key: "name", Value: "cmy"}}).Build())
}

func Test_logicalQueryBuilder_Nor(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$nor", Value: []any{bson.D{bson.E{Key: "name", Value: "cmy"}}}}}, NewBuilder().Nor(bson.D{{Key: "name", Value: "cmy"}}).Build())
}

func Test_logicalQueryBuilder_Or(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$or", Value: []any{bson.D{{Key: "name", Value: "cmy"}}}}}, NewBuilder().Or(bson.D{{Key: "name", Value: "cmy"}}).Build())
}
