package aggregation

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type accumulatorsBuilder struct {
	parent *Builder
}

func (b *accumulatorsBuilder) Sum(key string, expression any) *Builder {
	e := bson.E{Key: SumOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *accumulatorsBuilder) SumWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: SumOp, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Push(key string, expression any) *Builder {
	e := bson.E{Key: PushOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *accumulatorsBuilder) PushWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: PushOp, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Avg(key string, expression any) *Builder {
	e := bson.E{Key: AvgOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *accumulatorsBuilder) AvgWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: AvgOp, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) First(key string, expression any) *Builder {
	e := bson.E{Key: FirstOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *accumulatorsBuilder) FirstWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: FirstOp, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Last(key string, expression any) *Builder {
	e := bson.E{Key: LastOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *accumulatorsBuilder) LastWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: LastOp, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Min(key string, expression any) *Builder {
	e := bson.E{Key: MinOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *accumulatorsBuilder) MinWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: MinOp, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Max(key string, expression any) *Builder {
	e := bson.E{Key: MaxOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *accumulatorsBuilder) MaxWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: MaxOp, Value: expression})
	return b.parent
}
