package hello

import (
	"context"
	"github.com/bufbuild/connect-go"
	v1 "github.com/twoism/iocexample/gen/hello/v1"
	"github.com/twoism/iocexample/services/hello/container"
)

type Service struct {
	ctn container.Container
}

func NewService(ctn container.Container) *Service {
	return &Service{ctn: ctn}
}

func (s *Service) SayHello(ctx context.Context, req *connect.Request[v1.SayHelloRequest]) (*connect.Response[v1.SayHelloResponse], error) {
	s.ctn.Logger().Print("saying hello!")

	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}
