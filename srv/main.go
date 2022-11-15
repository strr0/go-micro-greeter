package main

import (
	"context"
	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
	"greeter/proto"
)

type Say struct {
	//
}

func (s *Say) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
	)
	service.Init()
	_ = greeter.RegisterSayHandler(service.Server(), new(Say))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
