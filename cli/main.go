package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	"greeter/proto"
	"log"
)

func main() {
	service := micro.NewService()
	service.Init()
	cl := greeter.NewSayService("go.micro.srv.greeter", service.Client())
	resp, err := cl.Hello(context.Background(), &greeter.Request{Name: "john"})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp.Msg)
}
