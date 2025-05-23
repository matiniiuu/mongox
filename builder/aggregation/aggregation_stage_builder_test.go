package aggregation

import (
	"testing"

	"github.com/matiniiuu/mongox/bsonx"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func TestStageBuilder_AddFields(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  mongo.Pipeline
	}{
		{
			name:  "nil value",
			value: nil,
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$addFields", Value: nil}}},
		},
		{
			name: "bson value",
			value: bson.D{
				bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
				bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$addFields", Value: bson.D{
					bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
					bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
				}}},
			},
		},
		{
			name: "map value",
			value: map[string]any{
				"totalHomework": bson.D{bson.E{Key: "$sum", Value: "$homework"}},
				"totalQuiz":     bson.D{bson.E{Key: "$sum", Value: "$quiz"}},
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$addFields", Value: map[string]any{
					"totalHomework": bson.D{bson.E{Key: "$sum", Value: "$homework"}},
					"totalQuiz":     bson.D{bson.E{Key: "$sum", Value: "$quiz"}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().AddFields(tc.value).Build())
		})
	}
}

func TestStageBuilder_Set(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  mongo.Pipeline
	}{
		{
			name:  "nil value",
			value: nil,
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$set", Value: nil}}},
		},
		{
			name: "bson value",
			value: bson.D{
				bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
				bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$set", Value: bson.D{
					bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
					bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
				}}},
			},
		},
		{
			name: "map value",
			value: map[string]any{
				"totalHomework": bson.D{bson.E{Key: "$sum", Value: "$homework"}},
				"totalQuiz":     bson.D{bson.E{Key: "$sum", Value: "$quiz"}},
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$set", Value: map[string]any{
					"totalHomework": bson.D{bson.E{Key: "$sum", Value: "$homework"}},
					"totalQuiz":     bson.D{bson.E{Key: "$sum", Value: "$quiz"}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Set(tc.value).Build())
		})
	}
}

func TestStageBuilder_Bucket(t *testing.T) {
	testCases := []struct {
		name       string
		groupBy    any
		opt        *BucketOptions
		boundaries []any
		want       mongo.Pipeline
	}{
		{
			name:       "defaultKey and output are nil",
			groupBy:    "$year_born",
			opt:        nil,
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
				}}},
			},
		},
		{
			name:    "output is nil",
			groupBy: "$year_born",
			opt: &BucketOptions{
				DefaultKey: "Other",
				Output:     nil,
			},
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
					bson.E{Key: "default", Value: "Other"},
				}}},
			},
		},
		{
			name:    "defaultKey is empty",
			groupBy: "$year_born",
			opt: &BucketOptions{
				DefaultKey: nil,
				Output:     NewBuilder().Sum("count", 1).Build(),
			},
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
				}}},
			},
		},
		{
			name:    "boundaries is nil",
			groupBy: "$year_born",
			opt: &BucketOptions{
				DefaultKey: "Other",
				Output:     NewBuilder().Sum("count", 1).Build(),
			},
			boundaries: nil,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: ([]any)(nil)},
					bson.E{Key: "default", Value: "Other"},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
				}}},
			},
		},
		{
			name:    "all not nil",
			groupBy: "$year_born",
			opt: &BucketOptions{
				DefaultKey: "Other",
				Output:     NewBuilder().Sum("count", 1).Push("artists", NewBuilder().Concat("name", "$first_name", " ", "$last_name").KeyValue("year_born", "$year_born").Build()).Build(),
			},
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
					bson.E{Key: "default", Value: "Other"},
					bson.E{Key: "output", Value: bson.D{
						bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
						bson.E{Key: "artists", Value: bson.D{
							bson.E{Key: "$push", Value: bson.D{
								bson.E{Key: "name", Value: bson.D{bson.E{Key: "$concat", Value: []any{"$first_name", " ", "$last_name"}}}},
								bson.E{Key: "year_born", Value: "$year_born"},
							}},
						}},
					}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Bucket(tc.groupBy, tc.boundaries, tc.opt).Build())
		})
	}
}

func TestStageBuilder_BucketAuto(t *testing.T) {
	testCases := []struct {
		name    string
		groupBy any
		opt     *BucketAutoOptions
		buckets int
		want    mongo.Pipeline
	}{
		{
			name:    "output and granularity are nil",
			groupBy: "$price",
			opt:     nil,
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
				}}},
			},
		},
		{
			name:    "output is nil",
			groupBy: "$price",
			opt: &BucketAutoOptions{
				Granularity: "R5",
			},
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
					bson.E{Key: "granularity", Value: "R5"},
				}}},
			},
		},
		{
			name:    "granularity is empty",
			groupBy: "$price",
			opt: &BucketAutoOptions{
				Output: NewBuilder().Sum("count", 1).Build(),
			},
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
				}}},
			},
		},
		{
			name:    "normal",
			groupBy: "$price",
			opt: &BucketAutoOptions{
				Output:      NewBuilder().Sum("count", 1).Build(),
				Granularity: "R5",
			},
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
					bson.E{Key: "granularity", Value: "R5"},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().BucketAuto(tc.groupBy, tc.buckets, tc.opt).Build())
		})
	}
}

func TestStageBuilder_Match(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       mongo.Pipeline
	}{
		{
			name:       "expression is nil",
			expression: nil,
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$match", Value: nil}}},
		},
		{
			name:       "expression is not nil",
			expression: NewBuilder().KeyValue("author", "dave").Build(),
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "author", Value: "dave"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Match(tc.expression).Build())
		})
	}
}

func TestStageBuilder_Group(t *testing.T) {
	testCases := []struct {
		name         string
		id           any
		accumulators []bson.E
		want         mongo.Pipeline
	}{
		{
			name:         "id and accumulators are nil",
			id:           nil,
			accumulators: nil,
			want:         mongo.Pipeline{bson.D{bson.E{Key: "$group", Value: bson.D{bson.E{Key: "_id", Value: nil}}}}},
		},
		{
			name: "id is nil",
			id:   nil,
			accumulators: NewBuilder().
				Sum("totalSaleAmount", bsonx.D("$multiply", []any{"$price", "$quantity"})).
				Avg("averageQuantity", "$quantity").
				Sum("count", 1).Build(),
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$group", Value: bson.D{
					bson.E{Key: "_id", Value: nil},
					bson.E{Key: "totalSaleAmount", Value: bson.D{bson.E{Key: "$sum", Value: bson.D{bson.E{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}}},
					bson.E{Key: "averageQuantity", Value: bson.D{bson.E{Key: "$avg", Value: "$quantity"}}},
					bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
				}}},
			},
		},
		{
			name:         "accumulators is nil",
			id:           bsonx.D("x", "$x"),
			accumulators: nil,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$group", Value: bson.D{
					bson.E{Key: "_id", Value: bson.D{bson.E{Key: "x", Value: "$x"}}}}}},
			},
		},
		{
			name: "id and accumulators are not nil",
			id:   bsonx.D("x", "$x"),
			accumulators: NewBuilder().
				Sum("totalSaleAmount", bsonx.D("$multiply", []any{"$price", "$quantity"})).
				Avg("averageQuantity", "$quantity").
				Sum("count", 1).Build(),
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$group", Value: bson.D{
					bson.E{Key: "_id", Value: bson.D{bson.E{Key: "x", Value: "$x"}}},
					bson.E{Key: "totalSaleAmount", Value: bson.D{bson.E{Key: "$sum", Value: bson.D{bson.E{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}}},
					bson.E{Key: "averageQuantity", Value: bson.D{bson.E{Key: "$avg", Value: "$quantity"}}},
					bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Group(tc.id, tc.accumulators...).Build())
		})
	}
}

func TestStageBuilder_Sort(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  mongo.Pipeline
	}{
		{
			name:  "nil value",
			value: nil,
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$sort", Value: nil}}},
		},
		{
			name:  "bson value",
			value: bson.D{bson.E{Key: "name", Value: 1}, bson.E{Key: "age", Value: -1}},
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "name", Value: 1}, bson.E{Key: "age", Value: -1}}}}},
		},
		{
			name:  "map value",
			value: map[string]any{"name": 1, "age": -1},
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$sort", Value: map[string]any{"name": 1, "age": -1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Sort(tc.value).Build())
		})
	}
}

func TestStageBuilder_Project(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  mongo.Pipeline
	}{
		{
			name:  "nil value",
			value: nil,
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$project", Value: nil}}},
		},
		{
			name:  "bson value",
			value: bson.D{bson.E{Key: "title", Value: 1}, bson.E{Key: "author", Value: 1}},
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$project", Value: bson.D{bson.E{Key: "title", Value: 1}, bson.E{Key: "author", Value: 1}}}}},
		},
		{
			name:  "map value",
			value: map[string]any{"title": 1, "author": 1},
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$project", Value: map[string]any{"title": 1, "author": 1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Project(tc.value).Build())
		})
	}
}

func TestStageBuilder_Limit(t *testing.T) {
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$limit", Value: int64(10)}}}, NewStageBuilder().Limit(10).Build())
}

func TestStageBuilder_Skip(t *testing.T) {
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$skip", Value: int64(10)}}}, NewStageBuilder().Skip(10).Build())
}

func TestStageBuilder_Unwind(t *testing.T) {
	testCases := []struct {
		name string
		path string
		opt  *UnWindOptions
		want mongo.Pipeline
	}{
		{
			name: "opt is nil",
			path: "$sizes",
			opt:  nil,
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: "$sizes"}}},
		},
		{
			name: "opt is not nil and includeArrayIndex is not empty",
			path: "$sizes",
			opt: &UnWindOptions{
				IncludeArrayIndex: "arrayIndex",
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: bson.D{
				bson.E{Key: "path", Value: "$sizes"},
				bson.E{Key: "includeArrayIndex", Value: "arrayIndex"},
			},
			}}},
		},
		{
			name: "opt is not nil and preserveNullAndEmptyArrays is true",
			path: "$sizes",
			opt: &UnWindOptions{
				PreserveNullAndEmptyArrays: true,
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: bson.D{
				bson.E{Key: "path", Value: "$sizes"},
				bson.E{Key: "preserveNullAndEmptyArrays", Value: true},
			}}}},
		},
		{
			name: "opt is not nil and includeArrayIndex is not empty and preserveNullAndEmptyArrays is true",
			path: "$sizes",
			opt: &UnWindOptions{
				IncludeArrayIndex:          "arrayIndex",
				PreserveNullAndEmptyArrays: true,
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: bson.D{
				bson.E{Key: "path", Value: "$sizes"},
				bson.E{Key: "includeArrayIndex", Value: "arrayIndex"},
				bson.E{Key: "preserveNullAndEmptyArrays", Value: true},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Unwind(tc.path, tc.opt).Build())
		})
	}
}

func TestStageBuilder_ReplaceRoot(t *testing.T) {
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "cities"}}}}}, NewStageBuilder().ReplaceRoot("cities").Build())
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: bson.D{{Key: "$mergeObjects", Value: []string{"cities", "countries"}}}}}}}}, NewStageBuilder().ReplaceRoot(bson.D{{Key: "$mergeObjects", Value: []string{"cities", "countries"}}}).Build())
}

func TestStageBuilder_ReplaceWith(t *testing.T) {
	testCases := []struct {
		name                string
		replacementDocument any
		want                mongo.Pipeline
	}{
		{
			name:                "nil replacementDocument",
			replacementDocument: nil,
			want:                mongo.Pipeline{bson.D{bson.E{Key: "$replaceWith", Value: nil}}},
		},
		{
			name:                "replacementDocument of string",
			replacementDocument: "$name",
			want:                mongo.Pipeline{bson.D{bson.E{Key: "$replaceWith", Value: "$name"}}},
		},
		//{
		//	name:                "replacementDocument of bson.D",
		//	replacementDocument: NewBuilder().ArrayToObject("$items").Build(),
		//	want:                mongo.Pipeline{bson.D{bson.E{Key: "$replaceWith", Value: bson.D{bson.E{Key: "$arrayToObject", Value: "$items"}}}}},
		//},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().ReplaceWith(tc.replacementDocument).Build())
		})
	}
}

func TestStageBuilder_Facet(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  mongo.Pipeline
	}{
		{
			name:  "nil facets",
			value: nil,
			want:  mongo.Pipeline{bson.D{bson.E{Key: "$facet", Value: nil}}},
		},
		//{
		//	// [
		//	//  {
		//	//    $facet: {
		//	//      "categorizedByTags": [
		//	//        { $unwind: "$tags" },
		//	//        { $sortByCount: "$tags" }
		//	//      ],
		//	//      "categorizedByPrice": [
		//	//        // Filter out documents without a price e.g., _id: 7
		//	//        { $match: { price: { $exists: 1 } } },
		//	//        {
		//	//          $bucket: {
		//	//            groupBy: "$price",
		//	//            boundaries: [  0, 150, 200, 300, 400 ],
		//	//            default: "Other",
		//	//            output: {
		//	//              "count": { $sum: 1 },
		//	//              "titles": { $push: "$title" }
		//	//            }
		//	//          }
		//	//        }
		//	//      ],
		//	//      "categorizedByYears(Auto)": [
		//	//        {
		//	//          $bucketAuto: {
		//	//            groupBy: "$year",
		//	//            buckets: 4
		//	//          }
		//	//        }
		//	//      ]
		//	//    }
		//	//  }
		//	//]
		//	name: "replacementDocument of bson.D",
		//	value: bsonx.NewD().
		//		Add("categorizedByTags", StageNewBuilder().Unwind("$tags", nil).SortByCount("$tags").Build()).
		//		Add("categorizedByPrice", StageNewBuilder().Match(
		//			NewBuilder().KeyValue("price", NewBuilder().KeyValue("$exists", 1).Build()).Build()).Bucket("$price", []any{0, 150, 200, 300, 400}, &BucketOptions{
		//			DefaultKey: "Other",
		//			Output:     NewBuilder().KeyValue("count", NewBuilder().Sum(1).Build()).KeyValue("titles", NewBuilder().Push("$title").Build()).Build(),
		//		}).Build()).
		//		Add("categorizedByYears(Auto)", StageNewBuilder().BucketAuto("$year", 4, nil).Build()).Build(),
		//	want: mongo.Pipeline{bson.D{bson.E{Key: "$facet", Value: bson.D{
		//		bson.E{Key: "categorizedByTags", Value: mongo.Pipeline{
		//			bson.D{bson.E{Key: "$unwind", Value: "$tags"}},
		//			bson.D{bson.E{Key: "$sortByCount", Value: "$tags"}},
		//		}},
		//		bson.E{Key: "categorizedByPrice", Value: mongo.Pipeline{
		//			bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "price", Value: bson.D{bson.E{Key: "$exists", Value: 1}}}}}},
		//			bson.D{bson.E{Key: "$bucket", Value: bson.D{
		//				bson.E{Key: "groupBy", Value: "$price"},
		//				bson.E{Key: "boundaries", Value: []any{0, 150, 200, 300, 400}},
		//				bson.E{Key: "default", Value: "Other"},
		//				bson.E{Key: "output", Value: bson.D{
		//					bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
		//					bson.E{Key: "titles", Value: bson.D{bson.E{Key: "$push", Value: "$title"}}},
		//				}},
		//			},
		//			}},
		//		}},
		//		bson.E{Key: "categorizedByYears(Auto)", Value: mongo.Pipeline{
		//			bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
		//				bson.E{Key: "groupBy", Value: "$year"},
		//				bson.E{Key: "buckets", Value: 4},
		//			}}},
		//		}},
		//	}}}},
		//},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Facet(tc.value).Build())
		})
	}
}

func TestStageBuilder_SortByCount(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       mongo.Pipeline
	}{
		{
			name:       "nil expression",
			expression: nil,
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$sortByCount", Value: nil}}},
		},
		{
			name:       "expression of string",
			expression: "$tags",
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$sortByCount", Value: "$tags"}}},
		},
		// { $sortByCount: { lname: "$employee.last", fname: "$employee.first" } }
		{
			name:       "expression of bson.D",
			expression: NewBuilder().KeyValue("lname", "$employee.last").KeyValue("fname", "$employee.first").Build(),
			want: mongo.Pipeline{bson.D{bson.E{Key: "$sortByCount", Value: bson.D{
				bson.E{Key: "lname", Value: "$employee.last"},
				bson.E{Key: "fname", Value: "$employee.first"},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().SortByCount(tc.expression).Build())
		})
	}
}

func TestStageBuilder_Count(t *testing.T) {
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$count", Value: "passing_scores"}}}, NewStageBuilder().Count("passing_scores").Build())
}

func TestStageBuilder_Lookup(t *testing.T) {
	testCases := []struct {
		name string
		from string
		as   string
		opt  *LookUpOptions

		want mongo.Pipeline
	}{
		{
			name: "basic",
			from: "orders",
			opt: &LookUpOptions{
				LocalField:   "_id",
				ForeignField: "userId",
				Let:          nil,
				Pipeline:     nil,
			},
			as: "userOrders",
			want: mongo.Pipeline{
				{
					bson.E{Key: "$lookup", Value: bson.D{
						bson.E{Key: "from", Value: "orders"},
						bson.E{Key: "localField", Value: "_id"},
						bson.E{Key: "foreignField", Value: "userId"},
						bson.E{Key: "as", Value: "userOrders"},
					}},
				},
			},
		},
		{
			name: "advanced case",
			from: "orders",
			opt: &LookUpOptions{
				LocalField:   "",
				ForeignField: "",
				Let:          bson.D{bson.E{Key: "userId", Value: "$_id"}},
				Pipeline: mongo.Pipeline{
					{
						bson.E{Key: "$match", Value: bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$and", Value: []any{
							bson.D{bson.E{Key: "$eq", Value: []any{"$userId", "$$userId"}}},
							bson.D{bson.E{Key: "$gt", Value: []any{"$totalAmount", 100}}},
						}}}}}},
					},
				},
			},
			as: "largeOrders",
			want: mongo.Pipeline{
				{
					bson.E{Key: "$lookup", Value: bson.D{
						bson.E{Key: "from", Value: "orders"},
						bson.E{Key: "let", Value: bson.D{bson.E{Key: "userId", Value: "$_id"}}},
						bson.E{Key: "pipeline", Value: mongo.Pipeline{
							{
								bson.E{Key: "$match", Value: bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$and", Value: []any{
									bson.D{bson.E{Key: "$eq", Value: []any{"$userId", "$$userId"}}},
									bson.D{bson.E{Key: "$gt", Value: []any{"$totalAmount", 100}}},
								}}}}}},
							},
						}},
						bson.E{Key: "as", Value: "largeOrders"},
					},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().Lookup(tc.from, tc.as, tc.opt).Build())
		})
	}
}

func TestStageBuilder_UnionWith(t *testing.T) {
	testCases := []struct {
		name     string
		coll     string
		pipeline mongo.Pipeline

		want mongo.Pipeline
	}{
		{
			name:     "basic",
			coll:     "orders",
			pipeline: mongo.Pipeline{},
			want: mongo.Pipeline{
				{
					bson.E{Key: "$unionWith", Value: "orders"},
				},
			},
		},
		{
			name: "advanced case",
			coll: "orders",
			pipeline: mongo.Pipeline{
				{
					bson.E{Key: "$match", Value: bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$and", Value: []any{
						bson.D{bson.E{Key: "$eq", Value: []any{"$userId", "$$userId"}}},
						bson.D{bson.E{Key: "$gt", Value: []any{"$totalAmount", 100}}},
					}}}}}},
				},
			},
			want: mongo.Pipeline{
				{
					bson.E{Key: "$unionWith", Value: bson.D{
						bson.E{Key: "coll", Value: "orders"},
						bson.E{Key: "pipeline", Value: mongo.Pipeline{
							{
								bson.E{Key: "$match", Value: bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$and", Value: []any{
									bson.D{bson.E{Key: "$eq", Value: []any{"$userId", "$$userId"}}},
									bson.D{bson.E{Key: "$gt", Value: []any{"$totalAmount", 100}}},
								}}}}}},
							},
						}},
					},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewStageBuilder().UnionWith(tc.coll, tc.pipeline).Build())
		})
	}
}
