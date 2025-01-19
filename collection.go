package mongox

import (
	"github.com/matiniiuu/mongox/aggregator"
	"github.com/matiniiuu/mongox/creator"
	"github.com/matiniiuu/mongox/deleter"
	"github.com/matiniiuu/mongox/finder"
	"github.com/matiniiuu/mongox/updater"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewCollection[T any](collection *mongo.Collection) *Collection[T] {
	return &Collection[T]{collection: collection}
}

type Collection[T any] struct {
	collection *mongo.Collection
}

func (c *Collection[T]) Finder() *finder.Finder[T] {
	return finder.NewFinder[T](c.collection)
}

func (c *Collection[T]) Creator() *creator.Creator[T] {
	return creator.NewCreator[T](c.collection)
}

func (c *Collection[T]) Updater() *updater.Updater[T] {
	return updater.NewUpdater[T](c.collection)
}

func (c *Collection[T]) Deleter() *deleter.Deleter[T] {
	return deleter.NewDeleter[T](c.collection)
}
func (c *Collection[T]) Aggregator() *aggregator.Aggregator[T] {
	return aggregator.NewAggregator[T](c.collection)
}

func (c *Collection[T]) Collection() *mongo.Collection {
	return c.collection
}
