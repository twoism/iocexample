package main

import (
	"github.com/rodaine/ioc"
	"log"
)

type ContainerOption func(container *ioc.Container)

type Container interface {
	Name() string
	Logger() *log.Logger
}

type container struct {
	c *ioc.Container

	name string
}

func NewDefaultContainer() Container {
	return NewContainer("example", WithLogger(log.Default()))
}

func NewContainer(name string, opts ...ContainerOption) Container {
	ctn := new(ioc.Container)

	for _, opt := range opts {
		opt(ctn)
	}
	ctn.Freeze()
	return &container{
		name: name,
		c:    ctn,
	}
}

func (c *container) Name() string {
	return c.name
}

func (c *container) Logger() *log.Logger {
	return ioc.Resolve[*log.Logger](c.c)
}

func WithLogger(l *log.Logger) ContainerOption {
	return func(ctn *ioc.Container) {
		ioc.Bind(ctn, ioc.Static[*log.Logger](l))
	}
}
