package main

import (
	"fmt"
	"github.com/twoism/iocexample/gen/hello/v1/hellov1connect"
	"github.com/twoism/iocexample/services/hello"
	"github.com/twoism/iocexample/services/hello/container"
	"net/http"
)

const addr = ":8080"

func main() {
	svc := hello.NewService(container.NewDefault())
	mux := http.NewServeMux()
	path, handler := hellov1connect.NewHelloServiceHandler(svc)
	mux.Handle(path, handler)
	fmt.Println("listening on", addr)
	http.ListenAndServe(addr, mux)
}
