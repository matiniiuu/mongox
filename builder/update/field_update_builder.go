package update

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type fieldUpdateBuilder struct {
	parent *Builder
}

func (b *fieldUpdateBuilder) Set(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(SetOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: SetOp, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Unset(keys ...string) *Builder {
	value := bson.D{}
	for i := range keys {
		value = append(value, bson.E{Key: keys[i], Value: ""})
	}
	b.parent.data = append(b.parent.data, bson.E{Key: UnsetOp, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) SetOnInsert(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(SetOnInsertOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: SetOnInsertOp, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) CurrentDate(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(CurrentDateOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: CurrentDateOp, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Inc(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(IncOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: IncOp, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Min(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(MinOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: MinOp, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Max(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(MaxOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: MaxOp, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Mul(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(MulOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: MulOp, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Rename(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(RenameOp, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: RenameOp, Value: bson.D{e}})
	}
	return b.parent
}
