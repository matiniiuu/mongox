package mongox

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Base struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt *time.Time    `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt *time.Time    `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DeletedAt *time.Time    `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

func (m *Base) DefaultId() bson.ObjectID {
	if m.ID.IsZero() {
		m.ID = bson.NewObjectID()
	}
	return m.ID
}

func (m *Base) DefaultCreatedAt() *time.Time {
	if m.CreatedAt == nil || m.CreatedAt.IsZero() {
		now := time.Now().Local()
		m.CreatedAt = &now
	}
	return m.CreatedAt
}

func (m *Base) DefaultUpdatedAt() *time.Time {
	now := time.Now().Local()
	m.UpdatedAt = &now
	return m.UpdatedAt
}
