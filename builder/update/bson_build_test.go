package update

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestAddToSet(t *testing.T) {
	t.Run("test AddToSet", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "colors", Value: "mauve"}}}}, AddToSet("colors", "mauve"))
	})
}

func TestPop(t *testing.T) {
	t.Run("test Pop", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$pop", Value: bson.D{bson.E{Key: "scores", Value: 1}}}}, Pop("scores", 1))
	})
}

func TestPull(t *testing.T) {
	t.Run("test Pull", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$pull", Value: bson.D{bson.E{Key: "fruit", Value: "apples"}}}}, Pull("fruit", "apples"))
	})
}

func TestPush(t *testing.T) {
	t.Run("test Push", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "scores", Value: 89}}}}, Push("scores", 89))
	})
}

func TestPullAll(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name:   "test PullAll",
			key:    "letters",
			values: []any{"b", "c"},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "letters", Value: []any{"b", "c"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, PullAll(tc.key, tc.values...))
		})
	}
}

func TestEach(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name:   "test Each",
			values: []any{3, 4, 5},
			want:   bson.D{bson.E{Key: "$each", Value: []any{3, 4, 5}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Each(tc.values...))
		})
	}
}

func TestPosition(t *testing.T) {
	t.Run("test Position", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$position", Value: 0}}, Position(0))
	})
}

func TestSlice(t *testing.T) {
	t.Run("test Slice", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$slice", Value: 3}}, Slice(3))
	})
}

func TestSort(t *testing.T) {
	t.Run("test Sort", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$sort", Value: 1}}, Sort(1))
	})
}

func TestSet(t *testing.T) {
	t.Run("test Set", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "Alice"}}}}, Set("name", "Alice"))
	})
}

func TestUnset(t *testing.T) {
	testCases := []struct {
		name  string
		value []string
		want  bson.D
	}{
		{
			name:  "test Unset",
			value: []string{"name", "age"},
			want:  bson.D{bson.E{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}, bson.E{Key: "age", Value: ""}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Unset(tc.value...))
		})
	}
}

func TestSetOnInsert(t *testing.T) {
	t.Run("test SetOnInsert", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "Alice"}}}}, SetOnInsert("name", "Alice"))
	})
}

func TestCurrentDate(t *testing.T) {
	t.Run("test CurrentDate", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}}}}, CurrentDate("lastModified", true))
	})
}

func TestInc(t *testing.T) {
	t.Run("test Inc", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "count", Value: 1}}}}, Inc("count", 1))
	})
}

func TestMin(t *testing.T) {
	t.Run("test Min", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$min", Value: bson.D{bson.E{Key: "lowScore", Value: 200}}}}, Min("lowScore", 200))
	})
}

func TestMax(t *testing.T) {
	t.Run("test Max", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$max", Value: bson.D{bson.E{Key: "highScore", Value: 800}}}}, Max("highScore", 800))
	})
}

func TestMul(t *testing.T) {
	t.Run("test Mul", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "qty", Value: 2}}}}, Mul("qty", 2))
	})
}

func TestRename(t *testing.T) {
	t.Run("test Rename", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$rename", Value: bson.D{bson.E{Key: "nickname", Value: "alias"}}}}, Rename("nickname", "alias"))
	})
}
