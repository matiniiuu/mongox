package model

import "context"

type BeforeInsert interface {
	BeforeInsert(ctx context.Context) error
}

type AfterInsert interface {
	AfterInsert(ctx context.Context) error
}

type BeforeUpdate interface {
	BeforeUpdate(ctx context.Context) error
}

type AfterUpdate interface {
	AfterUpdate(ctx context.Context) error
}

type BeforeUpsert interface {
	BeforeUpsert(ctx context.Context) error
}

type AfterUpsert interface {
	AfterUpsert(ctx context.Context) error
}

type BeforeDelete interface {
	BeforeDelete(ctx context.Context) error
}

type AfterDelete interface {
	AfterDelete(ctx context.Context) error
}

type BeforeFind interface {
	BeforeFind(ctx context.Context) error
}

type AfterFind interface {
	AfterFind(ctx context.Context) error
}
