package query

import (
	"github.com/matiniiuu/mongox/internal/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type evaluationQueryBuilder struct {
	parent *Builder
}

func (b *evaluationQueryBuilder) Expr(d bson.D) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: ExprOp, Value: d})
	return b.parent
}

func (b *evaluationQueryBuilder) JsonSchema(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: JsonSchemaOp, Value: value})
	return b.parent
}

func (b *evaluationQueryBuilder) Mod(key string, divisor any, remainder int) *Builder {
	if utils.IsNumeric(divisor) {
		e := bson.E{Key: ModOp, Value: bson.A{divisor, remainder}}
		if !b.parent.tryMergeValue(key, e) {
			b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
		}
	}
	return b.parent
}

func (b *evaluationQueryBuilder) Regex(key, value string) *Builder {
	e := bson.E{Key: RegexOp, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *evaluationQueryBuilder) RegexOptions(key, value, options string) *Builder {
	if !b.parent.tryMergeValue(key, bson.E{Key: RegexOp, Value: value}, bson.E{Key: OptionsOp, Value: options}) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{bson.E{Key: RegexOp, Value: value}, bson.E{Key: OptionsOp, Value: options}}})
	}
	return b.parent
}

// Text
// 如果 language 的值为零值，则不作为查询条件 If the value of language is zero, it is not used as a query condition
// 如果 caseSensitive 的值为零值，则不作为查询条件 If the value of caseSensitive is zero, it is not used as a query condition
// 如果 diacriticSensitive 的值为零值，则不作为查询条件 If the value of diacriticSensitive is zero, it is not used as a query condition
func (b *evaluationQueryBuilder) Text(value, language string, caseSensitive, diacriticSensitive bool) *Builder {
	d := bson.D{bson.E{Key: SearchOp, Value: value}}
	if language != "" {
		d = append(d, bson.E{Key: LanguageOp, Value: language})
	}
	if caseSensitive {
		d = append(d, bson.E{Key: CaseSensitiveOp, Value: caseSensitive})
	}
	if diacriticSensitive {
		d = append(d, bson.E{Key: DiacriticSensitiveOp, Value: diacriticSensitive})
	}
	b.parent.data = append(b.parent.data, bson.E{Key: TextOp, Value: d})
	return b.parent
}

func (b *evaluationQueryBuilder) Where(value string) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: WhereOp, Value: value})
	return b.parent
}
