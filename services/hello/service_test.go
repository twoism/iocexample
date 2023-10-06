package hello

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	hellov1 "github.com/twoism/iocexample/gen/hello/v1"
	"github.com/twoism/iocexample/services/hello/container"
	"testing"
)

func TestSayHello(t *testing.T) {
	ctn := container.New("test", container.WithFancyLogger())
	svc := NewService(ctn)

	resp, err := svc.SayHello(context.Background(),
		connect.NewRequest(&hellov1.SayHelloRequest{Name: "Chris"}))

	assert.Nil(t, resp)
	assert.Error(t, err)
}
