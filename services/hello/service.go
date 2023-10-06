package hello

import (
	"context"
	"github.com/bufbuild/connect-go"
	v1 "github.com/twoism/iocexample/gen/hello/v1"
	"github.com/twoism/iocexample/services/hello/container"
)

// Service is the hello service.
type Service struct {
	container.Container
}

// NewService returns a new hello service given a container with all its dependencies.
func NewService(ctn container.Container) *Service {
	return &Service{Container: ctn}
}

// SayHello implements the HelloService/SayHello method.
func (s *Service) SayHello(ctx context.Context, req *connect.Request[v1.SayHelloRequest]) (*connect.Response[v1.SayHelloResponse], error) {
	s.Logger().Print("saying hello!") // Use the injected logger.

	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}
