package main

import (
	"fmt"
	"github.com/twoism/iocexample/gen/hello/v1/hellov1connect"
	"net/http"
)

const addr = ":8080"

func main() {
	svc := NewService(NewDefaultContainer())
	mux := http.NewServeMux()
	path, handler := hellov1connect.NewHelloServiceHandler(svc)
	mux.Handle(path, handler)
	fmt.Println("listening on", addr)
	http.ListenAndServe(addr, mux)
}
