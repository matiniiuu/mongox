package operation

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OpType string

const (
	OpTypeBeforeInsert OpType = "beforeInsert"
	OpTypeAfterInsert  OpType = "afterInsert"
	OpTypeBeforeUpdate OpType = "beforeUpdate"
	OpTypeAfterUpdate  OpType = "afterUpdate"
	OpTypeBeforeDelete OpType = "beforeDelete"
	OpTypeAfterDelete  OpType = "afterDelete"
	OpTypeBeforeUpsert OpType = "beforeUpsert"
	OpTypeAfterUpsert  OpType = "afterUpsert"
	OpTypeBeforeFind   OpType = "beforeFind"
	OpTypeAfterFind    OpType = "afterFind"
)

type OpContext struct {
	Col *mongo.Collection `opt:"-"`
	Doc any
	// filter also can be used as query
	Filter       any
	Updates      any
	Replacement  any
	MongoOptions any
	ModelHook    any
}
