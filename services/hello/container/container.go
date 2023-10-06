package container

import (
	"github.com/rodaine/ioc"
	"log"
)

type Option func(container *ioc.Container)

type Container interface {
	Name() string
	Logger() *log.Logger
}

type container struct {
	c *ioc.Container

	name string
}

func NewDefault() Container {
	return New("example", WithLogger(log.Default()))
}

func New(name string, opts ...Option) Container {
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
