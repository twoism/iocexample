package container

import (
	"github.com/rodaine/ioc"
	"log"
	"os"
)

func WithLogger(l *log.Logger) Option {
	return func(ctn *ioc.Container) {
		ioc.Bind(ctn, ioc.Static[*log.Logger](l))
	}
}

func WithFancyLogger() Option {
	return WithLogger(
		log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	)
}
