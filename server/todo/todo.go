package todo

import (
	"context"
	"log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	tpb "grpc-todo/proto/todo"
)

type Server struct{}

type Todo struct{
	ID 			int64
	Title 		string
	Description string
}

var todos = []Todo{ 
	{
		ID: 1,
		Title:"This is title 1",
		Description:"This is description 1",
	},
	{
		ID: 2,
		Title:"This is title 2",
		Description:"This is description 2",
	},
	{
		ID: 3,
		Title:"This is title 3",
		Description:"This is description 3",
	},
}

func (s *Server) GetTodo(ctx context.Context, req *tpb.GetTodoRequest) (*tpb.GetTodoResponse, error){
	log.Printf("Todo ID: %d",req.GetID())

	var td Todo
	for _, value := range todos{
		if value.ID == req.GetID(){
			td = value
			break
		}
	}

	if td.ID == 0 {
		return &tpb.GetTodoResponse{},status.Errorf(codes.NotFound,"Invalid id")
	}
	return &tpb.GetTodoResponse{
		ID:td.ID,
		Title:td.Title,
		Description:td.Description,
	},nil
}

func (s *Server) GetTodos(req *tpb.GetTodosRequest, stream tpb.TodoService_GetTodosServer) error{
	
	for _, value := range todos{
		err := stream.Send(&tpb.GetTodoResponse{
			ID: value.ID,
			Title: value.Title,
			Description: value.Description,
		})

		if err != nil{
			return status.Error(codes.Internal,"Failed to send data" )
		}
	}

	return nil
}