package mongox

import (
	"testing"

	"github.com/matiniiuu/mongox/updater"

	"github.com/matiniiuu/mongox/creator"

	"github.com/matiniiuu/mongox/finder"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func TestCollection_New(t *testing.T) {
	mockMongoCollection := &mongo.Collection{}

	result := NewCollection[any](mockMongoCollection)

	assert.NotNil(t, result, "Expected non-nil Collection")
	assert.Equal(t, mockMongoCollection, result.collection, "Expected collection field to be initialized correctly")
}

func TestCollection_Finder(t *testing.T) {
	f := finder.NewFinder[any](&mongo.Collection{})
	assert.NotNil(t, f, "Expected non-nil Finder")
}

func TestCollection_Creator(t *testing.T) {
	c := creator.NewCreator[any](&mongo.Collection{})
	assert.NotNil(t, c, "Expected non-nil Creator")
}

func TestCollection_Updater(t *testing.T) {
	u := updater.NewUpdater[any](&mongo.Collection{})
	assert.NotNil(t, u, "Expected non-nil Updater")
}

func TestCollection_Deleter(t *testing.T) {
	d := NewCollection[any](&mongo.Collection{}).Deleter()
	assert.NotNil(t, d, "Expected non-nil Deleter")
}

func TestCollection_Aggregator(t *testing.T) {
	a := NewCollection[any](&mongo.Collection{}).Aggregator()
	assert.NotNil(t, a, "Expected non-nil Aggregator")
}

func TestCollection_Collection(t *testing.T) {
	a := NewCollection[any](&mongo.Collection{})
	assert.NotNil(t, a.Collection(), "Expected non-nil *mongo.Collection")
}
