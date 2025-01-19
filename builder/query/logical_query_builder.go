package query

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type logicalQueryBuilder struct {
	parent *Builder
}

// And
// 对于 conditions 参数，你同样可以使用 QueryBuilder 去生成
func (b *logicalQueryBuilder) And(conditions ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: AndOp, Value: conditions})
	return b.parent
}

func (b *logicalQueryBuilder) Not(condition any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: NotOp, Value: condition})
	return b.parent
}

// Nor
// 对于 conditions 参数，你同样可以使用 QueryBuilder 去生成
func (b *logicalQueryBuilder) Nor(conditions ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: NorOp, Value: conditions})
	return b.parent
}

// Or
// 对于 conditions 参数，你同样可以使用 QueryBuilder 去生成
func (b *logicalQueryBuilder) Or(conditions ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: OrOp, Value: conditions})
	return b.parent
}
