package hook

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DefaultModel interface {
	// DefaultId set and get
	DefaultId() bson.ObjectID
	// DefaultCreatedAt set and get
	DefaultCreatedAt() time.Time
	// DefaultUpdatedAt set and get
	DefaultUpdatedAt() time.Time
}

type CustomModel interface {
	// CustomID set and get field name and value
	CustomID() (string, any)
	// CustomCreatedAt set and get field name and value
	CustomCreatedAt() (string, any)
	// CustomUpdatedAt set and get field name and value
	CustomUpdatedAt() (string, any)
}
