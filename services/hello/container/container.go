package container

import (
	"github.com/rodaine/ioc"
	"log"
)

// Option is a function that configures a container.
type Option func(container *ioc.Container)

// Container is the interface for the hello service container.
// Accessors for dependencies should be exposed here.
type Container interface {
	Name() string
	Logger() *log.Logger
}

type container struct {
	c *ioc.Container

	name string
}

// NewDefault returns a new container with all its default dependencies.
func NewDefault() Container {
	return New("example", WithLogger(InitFancyLogger()))
}

// New returns an empty container.
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

// Name returns the name of the container.
func (c *container) Name() string {
	return c.name
}

// Logger returns the injected logger.
func (c *container) Logger() *log.Logger {
	return ioc.Resolve[*log.Logger](c.c)
}
