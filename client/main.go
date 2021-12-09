package main

import (
	"grpc-todo/client/todo"
	"log"

	"google.golang.org/grpc"
)

func main(){
		conn, err :=grpc.Dial(":3000",grpc.WithInsecure())

		if err !=nil{
			log.Fatalf("Failed to connect: %s", err)
		}

		c:=todo.NewClient(conn)
		
		res,err := c.GotTodo(3)

		if err !=nil{
			log.Fatalf("err while calling GetTodo: %s",err)
		}
		log.Printf("Response %#v", res)

		if err:=c.GetTodos(); err !=nil{
			log.Fatalf("err while calling GetTodos: %s",err)
		}

		
}