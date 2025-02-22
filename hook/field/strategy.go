package field

import (
	"github.com/matiniiuu/mongox/hook"
	"github.com/matiniiuu/mongox/operation"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	field struct {
		name  string
		value any
	}
)

var strategies = map[operation.OpType]func(doc any, opts ...any) error{
	operation.OpTypeBeforeInsert: BeforeInsert,
	operation.OpTypeBeforeUpdate: BeforeUpdate,
	operation.OpTypeBeforeUpsert: BeforeUpsert,
}

func BeforeInsert(doc any, _ ...any) error {
	if doc == nil {
		return nil
	}
	if defaultModel, ok := doc.(hook.DefaultModel); ok {
		defaultModel.DefaultId()
		defaultModel.DefaultCreatedAt()
	}

	if customModel, ok := doc.(hook.CustomModel); ok {
		customModel.CustomID()
		customModel.CustomCreatedAt()
	}

	return nil
}

func BeforeUpdate(doc any, opts ...any) error {
	if doc == nil {
		return nil
	}
	var (
		defaultModel hook.DefaultModel
		customModel  hook.CustomModel
		ok           bool
	)
	defaultModel, ok = doc.(hook.DefaultModel)
	if !ok {
		customModel, ok = doc.(hook.CustomModel)
		if !ok {
			return nil
		}
	}

	if len(opts) == 0 || opts[0] == nil {
		return nil
	}

	updates, ok := opts[0].(bson.M)
	if !ok || updates == nil {
		return nil
	}

	if updates["$set"] == nil {
		updates["$set"] = bson.M{}
	}

	setFields, ok := updates["$set"].(bson.M)
	if !ok || setFields == nil {
		return nil
	}

	updatedAtField := getField("updated_at", defaultModel, customModel)

	setFields[updatedAtField.name] = updatedAtField.value

	return nil
}

func BeforeUpsert(doc any, opts ...any) error {
	if doc == nil {
		return nil
	}
	var (
		defaultModel hook.DefaultModel
		customModel  hook.CustomModel
		ok           bool
	)
	defaultModel, ok = doc.(hook.DefaultModel)
	if !ok {
		customModel, ok = doc.(hook.CustomModel)
		if !ok {
			return nil
		}
	}

	if len(opts) == 0 || opts[0] == nil {
		return nil
	}

	updates, ok := opts[0].(bson.M)
	if !ok || updates == nil {
		return nil
	}

	if updates["$set"] == nil {
		updates["$set"] = bson.M{}
	}

	setFields, ok := updates["$set"].(bson.M)
	if !ok || setFields == nil {
		return nil
	}

	if updates["$setOnInsert"] == nil {
		updates["$setOnInsert"] = bson.M{}
	}

	setOnInsertFields, ok := updates["$setOnInsert"].(bson.M)
	if !ok || setOnInsertFields == nil {
		return nil
	}

	idField, createdAtField, updatedAtField := getField("_id", defaultModel, customModel), getField("created_at", defaultModel, customModel), getField("updated_at", defaultModel, customModel)
	setFields[updatedAtField.name] = updatedAtField.value

	setOnInsertFields[idField.name] = idField.value
	setOnInsertFields[createdAtField.name] = createdAtField.value

	return nil
}

func getField(filed string, defaultModel hook.DefaultModel, customModel hook.CustomModel) field {
	var (
		name  string
		value any
	)
	switch filed {
	case "_id":
		if defaultModel != nil {
			return field{name: "_id", value: defaultModel.DefaultId()}
		}
		name, value = customModel.CustomID()
	case "created_at":
		if defaultModel != nil {
			return field{name: "created_at", value: defaultModel.DefaultCreatedAt()}
		}
		name, value = customModel.CustomCreatedAt()
	case "updated_at":
		if defaultModel != nil {
			return field{name: "updated_at", value: defaultModel.DefaultUpdatedAt()}
		}
		name, value = customModel.CustomUpdatedAt()
	default:
		return field{name: "", value: nil}
	}
	return field{name: name, value: value}
}
