package todo

import (
	"context"
	"io"
	"log"


	tpb "grpc-todo/proto/todo"

	"google.golang.org/grpc"
)

type Client struct{
	Client tpb.TodoServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client{
	return Client{
		Client: tpb.NewTodoServiceClient(conn),
	}
}

func (c *Client) GotTodo(id int64)(*tpb.GetTodoResponse, error){
	return c.Client.GetTodo(context.Background(), &tpb.GetTodoRequest{
		ID: id,
	})

}

func (c *Client) GetTodos() error {
	stream,err :=c.Client.GetTodos(context.Background(),&tpb.GetTodosRequest{})
	if err !=nil{
		return err
	}

	for{
		res, err := stream.Recv()
		if err !=nil{
			if err == io.EOF{
				return nil
			}
			return err
		}
		log.Printf("Received Id: %d", res.ID)
	}

}