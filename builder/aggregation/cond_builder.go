package aggregation

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type condBuilder struct {
	parent *Builder
}

func (b condBuilder) Cond(key string, boolExpr, tureExpr, falseExpr any) *Builder {
	e := bson.E{Key: CondOp, Value: []any{boolExpr, tureExpr, falseExpr}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b condBuilder) CondWithoutKey(boolExpr, tureExpr, falseExpr any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: CondOp, Value: []any{boolExpr, tureExpr, falseExpr}})
	return b.parent
}

func (b condBuilder) IfNull(key string, expr, replacement any) *Builder {
	e := bson.E{Key: IfNullOp, Value: []any{expr, replacement}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b condBuilder) IfNullWithoutKey(expr, replacement any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: IfNullOp, Value: []any{expr, replacement}})
	return b.parent
}

func (b condBuilder) Switch(key string, cases []CaseThen, defaultCase any) *Builder {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{{Key: CaseOp, Value: caseThen.Case}, {Key: ThenOp, Value: caseThen.Then}})
	}
	e := bson.E{Key: SwitchOp, Value: bson.D{bson.E{Key: BranchesOp, Value: branches}, bson.E{Key: DefaultCaseOp, Value: defaultCase}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b condBuilder) SwitchWithoutKey(cases []CaseThen, defaultCase any) *Builder {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{bson.E{Key: CaseOp, Value: caseThen.Case}, {Key: ThenOp, Value: caseThen.Then}})
	}
	b.parent.d = append(b.parent.d, bson.E{Key: SwitchOp, Value: bson.D{bson.E{Key: BranchesOp, Value: branches}, bson.E{Key: DefaultCaseOp, Value: defaultCase}}})
	return b.parent
}
