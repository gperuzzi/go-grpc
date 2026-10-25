
package main

import (
	"context"
	"log"
	"time"

	tasksv1 "example.com/go-grpc/gen/tasks/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := tasksv1.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	// Create a task.
	task, err := client.CreateTask(ctx, &tasksv1.CreateTaskRequest{
		Title:       "Learn gRPC",
		Description: "Build my first Go gRPC service",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created task: %+v", task)

	// Retrieve the task using its ID.
	retrieved, err := client.GetTask(ctx, &tasksv1.GetTaskRequest{
		Id: task.GetId(),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Retrieved task: %+v", retrieved)
}
