package container

import (
	"github.com/rodaine/ioc"
	"github.com/twoism/iocexample/services/hello/internal"
)

func WithDB(db *internal.DB) Option {
	return func(ctn *ioc.Container) {
		ioc.Bind(ctn, ioc.Static[*internal.DB](db))
	}
}

func WithDefaultDB() *internal.DB {
	return internal.HelloDBConn()
}
