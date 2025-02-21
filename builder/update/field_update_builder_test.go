package update

import (
	"testing"
	"time"

	"github.com/matiniiuu/mongox/bsonx"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_fieldUpdateBuilder_Set(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}}, NewBuilder().Set("name", "cmy").Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}}, NewBuilder().Set("name", "cmy").Set("age", 24).Build())
	})
}

func Test_fieldUpdateBuilder_Unset(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{}}}, NewBuilder().Unset().Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}}}}, NewBuilder().Unset("name").Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}, bson.E{Key: "age", Value: ""}}}}, NewBuilder().Unset("name", "age").Build())
}

func Test_fieldUpdateBuilder_SetOnInsert(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}}, NewBuilder().SetOnInsert("name", "cmy").Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}}, NewBuilder().SetOnInsert("name", "cmy").SetOnInsert("age", 24).Build())
	})
}

func Test_fieldUpdateBuilder_Inc(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$inc", Value: bson.D{bson.E{Key: "orders", Value: 1}}}}, NewBuilder().Inc("orders", 1).Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$inc", Value: bson.D{bson.E{Key: "orders", Value: 1}, bson.E{Key: "ratings", Value: -1}}}}, NewBuilder().Inc("orders", 1).Inc("ratings", -1).Build())
	})
}

func Test_fieldUpdateBuilder_Min(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 100}}}}, NewBuilder().Min("stock", 100).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, NewBuilder().Min("stock", 100).Min("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Build())
	})
}

func Test_fieldUpdateBuilder_Max(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}}}}, NewBuilder().Max("stock", 100).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, NewBuilder().Max("stock", 100).Max("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Build())
	})
}

func Test_fieldUpdateBuilder_Mul(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}}}}, NewBuilder().Mul("price", 1.25).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "quantity", Value: 2}}}}, NewBuilder().Mul("price", 1.25).Mul("quantity", 2).Build())
	})
}

func Test_fieldUpdateBuilder_Rename(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$rename", Value: bson.D{bson.E{Key: "name", Value: "name"}}}}, NewBuilder().Rename("name", "name").Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$rename", Value: bson.D{bson.E{Key: "name", Value: "name"}, bson.E{Key: "age", Value: "age"}}}}, NewBuilder().Rename("name", "name").Rename("age", "age").Build())
	})
}

func Test_fieldUpdateBuilder_CurrentDate(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}}}}, NewBuilder().CurrentDate("lastModified", true).Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}, bson.E{Key: "cancellation.date", Value: bson.D{bson.E{Key: "$type", Value: "timestamp"}}}}}}, NewBuilder().CurrentDate("lastModified", true).CurrentDate("cancellation.date", bsonx.D("$type", "timestamp")).Build())
	})
}
