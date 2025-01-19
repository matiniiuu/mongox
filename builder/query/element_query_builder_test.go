package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_elementQueryBuilder_Exists(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$exists", Value: true}}}}, NewBuilder().Exists("name", true).Build())
}

func Test_elementQueryBuilder_Type(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$type", Value: bson.TypeString}}}}, NewBuilder().Type("name", bson.TypeString).Build())
}

func Test_elementQueryBuilder_TypeAlias(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$type", Value: "string"}}}}, NewBuilder().TypeAlias("name", "string").Build())
}

func TestNewBuilder_TypeArray(t *testing.T) {

	testCases := []struct {
		name string
		key  string
		ts   []bson.Type

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: ([]bson.Type)(nil)}}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []bson.Type{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []bson.Type{}}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []bson.Type{bson.TypeString},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []bson.Type{bson.TypeString}}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []bson.Type{bson.TypeString, bson.TypeInt32},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []bson.Type{bson.TypeString, bson.TypeInt32}}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBuilder().TypeArray(tc.key, tc.ts...).Build())
		})
	}
}

func TestNewBuilder_TypeArrayAlias(t *testing.T) {

	testCases := []struct {
		name string
		key  string
		ts   []string

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: ([]string)(nil)}}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []string{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []string{}}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []string{"string"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []string{"string"}}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []string{"string", "int32"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []string{"string", "int32"}}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBuilder().TypeArrayAlias(tc.key, tc.ts...).Build())
		})
	}
}
