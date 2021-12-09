package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	tpb "grpc-todo/proto/todo"
	"grpc-todo/server/todo"
)

func main(){
	lis,err :=net.Listen("tcp",":3000")
	if err != nil{

		log.Fatalf("Failed to Listen: %s", err)
	}

	grpcServer :=grpc.NewServer()

	s:=todo.Server{}

	tpb.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err!=nil{
		log.Fatalf("Failed to serve: %s", err)
	}
}