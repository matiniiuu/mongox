package aggregation

import (
	"testing"

	"github.com/matiniiuu/mongox/bsonx"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_condBuilder_Cond(t *testing.T) {
	t.Run("test Cond", func(t *testing.T) {
		assert.Equal(t,
			bson.D{bson.E{Key: "discount", Value: bson.D{{Key: "$cond", Value: []any{bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20}}}}},
			NewBuilder().Cond("discount", bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20).Build(),
		)
	})
}

func Test_condBuilder_CondWithoutKey(t *testing.T) {
	testCases := []struct {
		name      string
		boolExpr  any
		trueExpr  any
		falseExpr any
		expected  bson.D
	}{
		{
			name:      "normal",
			boolExpr:  NewBuilder().GteWithoutKey("$qty", 250).Build(),
			trueExpr:  30,
			falseExpr: 20,
			expected:  bson.D{{Key: "$cond", Value: []any{bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().CondWithoutKey(tc.boolExpr, tc.trueExpr, tc.falseExpr).Build())
		})
	}
}

func Test_condBuilder_IfNull(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "discount", Value: bson.D{{Key: "$ifNull", Value: []any{"$coupon", int64(0)}}}}}, NewBuilder().IfNull("discount", "$coupon", int64(0)).Build())
}

func Test_condBuilder_IfNullWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expr        any
		replacement any
		expected    bson.D
	}{
		{
			name:        "nil expr",
			expr:        nil,
			replacement: "Unspecified",
			expected:    bson.D{{Key: "$ifNull", Value: []any{nil, "Unspecified"}}},
		},
		{
			name:        "nil replacement",
			expr:        "$description",
			replacement: nil,
			expected:    bson.D{{Key: "$ifNull", Value: []any{"$description", nil}}},
		},
		{
			name:        "normal",
			expr:        "$description",
			replacement: "Unspecified",
			expected:    bson.D{{Key: "$ifNull", Value: []any{"$description", "Unspecified"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().IfNullWithoutKey(tc.expr, tc.replacement).Build())
		})
	}
}

func Test_condBuilder_Switch(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "summary", Value: bson.D{
		{Key: "$switch", Value: bson.D{
			{Key: "branches", Value: bson.A{
				bson.D{{Key: "case", Value: bson.D{{Key: "$eq", Value: []any{0, 5}}}}, {Key: "then", Value: "equals"}},
				bson.D{{Key: "case", Value: bson.D{{Key: "$gt", Value: []any{0, 5}}}}, {Key: "then", Value: "greater than"}},
			}},
			{Key: "default", Value: "Did not match"},
		}},
	}}},
		NewBuilder().Switch("summary", []CaseThen{
			{
				Case: bsonx.D("$eq", []any{0, 5}), Then: "equals",
			},
			{
				Case: bsonx.D("$gt", []any{0, 5}), Then: "greater than",
			},
		}, "Did not match").Build(),
	)
}

func Test_condBuilder_SwitchWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		cases       []CaseThen
		defaultCase any
		expected    bson.D
	}{
		{
			name:        "nil cases",
			cases:       nil,
			defaultCase: "Did not match",
			expected: bson.D{
				{Key: "$switch", Value: bson.D{
					{Key: "branches", Value: bson.A{}},
					{Key: "default", Value: "Did not match"},
				}},
			},
		},
		{
			name:        "empty cases",
			cases:       []CaseThen{},
			defaultCase: "Did not match",
			expected: bson.D{
				{Key: "$switch", Value: bson.D{
					{Key: "branches", Value: bson.A{}},
					{Key: "default", Value: "Did not match"},
				}},
			},
		},
		{
			name: "normal",
			cases: []CaseThen{
				{
					Case: bsonx.D("$eq", []any{0, 5}), Then: "equals",
				},
				{
					Case: bsonx.D("$gt", []any{0, 5}), Then: "greater than",
				},
			},
			defaultCase: "Did not match",
			expected: bson.D{
				{Key: "$switch", Value: bson.D{
					{Key: "branches", Value: bson.A{
						bson.D{{Key: "case", Value: bson.D{{Key: "$eq", Value: []any{0, 5}}}}, {Key: "then", Value: "equals"}},
						bson.D{{Key: "case", Value: bson.D{{Key: "$gt", Value: []any{0, 5}}}}, {Key: "then", Value: "greater than"}},
					}},
					{Key: "default", Value: "Did not match"},
				}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().SwitchWithoutKey(tc.cases, tc.defaultCase).Build())
		})
	}
}
