package container

import (
	"github.com/rodaine/ioc"
	"log"
)

func WithLogger(l *log.Logger) Option {
	return func(ctn *ioc.Container) {
		ioc.Bind(ctn, ioc.Static[*log.Logger](l))
	}
}
