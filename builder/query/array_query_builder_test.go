package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_arrayQueryBuilder_ElemMatch(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "age", Value: bson.D{bson.E{Key: "$elemMatch", Value: bson.D{bson.E{Key: "$gt", Value: 1}}}}}}, NewBuilder().ElemMatch("age", NewBuilder().KeyValue("$gt", 1).Build()).Build())
}

func TestNewBuilder_All(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []any

		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]any)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []any{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []any{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []any{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []any{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []any{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []any{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []any{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []any{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.builder.All(tc.key, tc.values...).Build())
		})
	}
}

func TestNewBuilder_AllUint(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []uint

		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]uint)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []uint{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []uint{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllUint8(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []uint8
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]uint8)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint8{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint8{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint8{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint8{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint8{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint8{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []uint8{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []uint8{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllUint16(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []uint16
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]uint16)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint16{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint16{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint16{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint16{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint16{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint16{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []uint16{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []uint16{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllUint32(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		builder *Builder
		values  []uint32

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]uint32)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint32{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint32{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint32{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint32{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint32{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint32{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []uint32{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []uint32{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllUint64(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []uint64
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]uint64)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint64{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint64{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint64{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint64{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []uint64{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []uint64{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []uint64{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []uint64{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllInt(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]int)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []int{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []int{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []int{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []int{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []int{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllInt(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllInt8(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int8
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]int8)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []int8{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int8{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []int8{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int8{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []int8{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int8{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []int8{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []int8{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := tc.builder.AllInt8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllInt16(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		values  []int16
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]int16)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []int16{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int16{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []int16{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int16{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []int16{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int16{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []int16{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []int16{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllInt16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllInt32(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int32
		builder *Builder
		want    bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]int32)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			values:  []int32{},
			builder: NewBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int32{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			values:  []int32{1},
			builder: NewBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int32{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			values:  []int32{1, 2, 3},
			builder: NewBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int32{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			values:  []int32{18, 19, 20},
			builder: NewBuilder().Gt("age", 18),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []int32{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := tc.builder.AllInt32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllInt64(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int64
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]int64)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			values:  []int64{},
			builder: NewBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int64{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			values:  []int64{1},
			builder: NewBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int64{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			values:  []int64{1, 2, 3},
			builder: NewBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []int64{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			values:  []int64{18, 19, 20},
			builder: NewBuilder().Gt("age", 18),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []int64{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllInt64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllFloat32(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []float32
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]float32)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []float32{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []float32{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []float32{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []float32{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []float32{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []float32{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []float32{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []float32{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllFloat32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllFloat64(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []float64
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]float64)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []float64{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []float64{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []float64{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []float64{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []float64{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []float64{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []float64{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []float64{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllFloat64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewBuilder_AllString(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []string
		builder *Builder
		want    bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: NewBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: ([]string)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: NewBuilder(),
			values:  []string{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []string{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: NewBuilder(),
			values:  []string{"1"},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []string{"1"}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: NewBuilder(),
			values:  []string{"1", "2", "3"},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: AllOp, Value: []string{"1", "2", "3"}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: NewBuilder().Gt("age", 18),
			values:  []string{"18", "19", "20"},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: AllOp, Value: []string{"18", "19", "20"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_arrayQueryBuilder_Size(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "age", Value: bson.D{bson.E{Key: "$size", Value: 1}}}}, NewBuilder().Size("age", 1).Build())

	assert.Equal(t, bson.D{{Key: "age", Value: bson.D{bson.E{Key: "$gt", Value: 18}, bson.E{Key: "$size", Value: 1}}}}, NewBuilder().Gt("age", 18).Size("age", 1).Build())
}
