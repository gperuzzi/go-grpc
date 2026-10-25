
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	tasksv1 "example.com/go-grpc/gen/tasks/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type taskServer struct {
	tasksv1.UnimplementedTaskServiceServer

	mu     sync.RWMutex
	tasks  map[string]*tasksv1.Task
	nextID int
}

func newTaskServer() *taskServer {
	return &taskServer{
		tasks: make(map[string]*tasksv1.Task),
	}
}

func (s *taskServer) CreateTask(
	ctx context.Context,
	req *tasksv1.CreateTaskRequest,
) (*tasksv1.Task, error) {
	if req.GetTitle() == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"title is required",
		)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	id := fmt.Sprintf("%d", s.nextID)

	task := &tasksv1.Task{
		Id:          id,
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	}

	s.tasks[id] = task
	log.Printf("Created task %s: %s", id, task.Title)

	return task, nil
}

func (s *taskServer) GetTask(
	ctx context.Context,
	req *tasksv1.GetTaskRequest,
) (*tasksv1.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.tasks[req.GetId()]
	if !ok {
		return nil, status.Error(
			codes.NotFound,
			"task not found",
		)
	}

	return task, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	tasksv1.RegisterTaskServiceServer(
		grpcServer,
		newTaskServer(),
	)

	log.Println("gRPC server listening on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
