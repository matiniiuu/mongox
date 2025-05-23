package aggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_comparisonBuilder_Eq(t *testing.T) {
	t.Run("test Eq", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{bson.E{Key: "$eq", Value: []any{"$qty", 250}}}}},
			NewBuilder().Eq("items", "$qty", 250).Build())
	})
}

func Test_comparisonBuilder_EqWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$eq", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$eq", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$eq", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().EqWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Ne(t *testing.T) {
	t.Run("test Ne", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{bson.E{Key: "$ne", Value: []any{"$qty", 250}}}}},
			NewBuilder().Ne("items", "$qty", 250).Build())
	})
}

func Test_comparisonBuilder_NeWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$ne", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$ne", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$ne", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().NeWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Gt(t *testing.T) {
	t.Run("test Gt", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}}},
			NewBuilder().Gt("items", "$qty", 250).Build())
	})
}

func Test_comparisonBuilder_GtWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$gt", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$gt", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().GtWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Gte(t *testing.T) {
	t.Run("test Gte", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{bson.E{Key: "$gte", Value: []any{"$qty", 250}}}}},
			NewBuilder().Gte("items", "$qty", 250).Build())
	})
}

func Test_comparisonBuilder_GteWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$gte", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$gte", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$gte", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().GteWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Lt(t *testing.T) {
	t.Run("test Lt", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}}}},
			NewBuilder().Lt("items", "$qty", 250).Build())
	})
}

func Test_comparisonBuilder_LtWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$lt", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$lt", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().LtWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Lte(t *testing.T) {
	t.Run("test Lte", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{bson.E{Key: "$lte", Value: []any{"$qty", 250}}}}},
			NewBuilder().Lte("items", "$qty", 250).Build())
	})
}

func Test_comparisonBuilder_LteWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$lte", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$lte", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$lte", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().LteWithoutKey(tc.expressions...).Build())
		})
	}
}
