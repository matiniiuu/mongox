package bsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestD(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "name", Value: "Mingyong Chen"}}, D("name", "Mingyong Chen"))
}

func TestId(t *testing.T) {
	assert.Equal(t, bson.M{"_id": "1"}, Id("1"))
}

func TestE(t *testing.T) {
	assert.Equal(t, bson.E{Key: "name", Value: "chenmingyong"}, E("name", "chenmingyong"))
}

func TestA(t *testing.T) {
	testCases := []struct {
		name   string
		values []any
		want   bson.A
	}{
		{
			name:   "nil values",
			values: nil,
			want:   bson.A{},
		},
		{
			name:   "empty values",
			values: []any{},
			want:   bson.A{},
		},
		{
			name:   "multiple values",
			values: []any{"1", "2"},
			want:   bson.A{"1", "2"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, A(tc.values...))
		})
	}

}
