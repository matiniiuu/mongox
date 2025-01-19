package query

const (
	AllOp                = "$all"
	AndOp                = "$and"
	CaseSensitiveOp      = "$caseSensitive"
	DiacriticSensitiveOp = "$diacriticSensitive"
	ElemMatchOp          = "$elemMatch"
	EqOp                 = "$eq"
	ExistsOp             = "$exists"
	ExprOp               = "$expr"
	GtOp                 = "$gt"
	GteOp                = "$gte"
	IdOp                 = "_id"
	InOp                 = "$in"
	JsonSchemaOp         = "$jsonSchema"
	LanguageOp           = "$language"
	LtOp                 = "$lt"
	LteOp                = "$lte"
	ModOp                = "$mod"
	NeOp                 = "$ne"
	NinOp                = "$nin"
	NorOp                = "$nor"
	NotOp                = "$not"
	OptionsOp            = "$options"
	OrOp                 = "$or"
	RegexOp              = "$regex"
	SearchOp             = "$search"
	SizeOp               = "$size"
	SliceOp              = "$slice"
	TextOp               = "$text"
	TypeOp               = "$type"
	WhereOp              = "$where"
)

type TextOptions struct {
	Language           string
	CaseSensitive      bool
	DiacriticSensitive bool
}
