package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	hellov1 "github.com/twoism/iocexample/gen/hello/v1"
	"log"
	"testing"
)

func TestSayHello(t *testing.T) {
	ctn := NewContainer("test", WithLogger(log.Default()))
	svc := NewService(ctn)

	resp, err := svc.SayHello(context.Background(),
		connect.NewRequest(&hellov1.SayHelloRequest{Name: "Chris"}))

	assert.Nil(t, resp)
	assert.Error(t, err)
}
