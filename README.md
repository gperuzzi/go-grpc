Go gRPC Task Service

A beginner-friendly gRPC application built with Go and Protocol Buffers. It demonstrates how a client communicates with a gRPC server to create and retrieve tasks.

Features
Create a task using gRPC.
Retrieve a task by ID.
Define the API using Protocol Buffers.
Generate Go code from .proto files.
Handle gRPC errors and concurrent requests.
Tech Stack
Go — application language
gRPC — remote procedure calls
Protocol Buffers — API and message definitions
Ubuntu on WSL 2 — development environment
VS Code — code editor
Project Structure
text
go-grpc/
├── client/
│   └── main.go
├── server/
│   └── main.go
├── proto/
│   └── tasks/
│       └── v1/
│           └── tasks.proto
├── gen/
│   └── tasks/
│       └── v1/
│           ├── tasks.pb.go
│           └── tasks_grpc.pb.go
├── go.mod
└── README.md

Prerequisites

Install the following tools:

Go
Protocol Buffers compiler
Go Protocol Buffers generator: protoc-gen-go
Go gRPC generator: protoc-gen-go-grpc
Installation

Clone the repository and enter the project directory:

bash
git clone git@github.com:gperuzzi/go-grpc.git
cd go-grpc


Install the Go dependencies:

bash
go mod tidy


If you modify the .proto definition, regenerate the Go code:

bash
protoc \
  --proto_path=. \
  --go_out=. \
  --go_opt=module=<your-module-path> \
  --go-grpc_out=. \
  --go-grpc_opt=module=<your-module-path> \
  proto/tasks/v1/tasks.proto


Replace <your-module-path> with the module path declared in go.mod.

Running the Application
1. Start the gRPC server

Open a terminal and run:

bash
go run ./server


The server listens on port 50051.

2. Run the client

Open a second terminal in the project directory:

bash
go run ./client


The client creates a task and retrieves it from the server.

Example output:

text
Created task: id:"1" title:"Learn gRPC"
Retrieved task: id:"1" title:"Learn gRPC"

Available RPC Methods
Method	Description
CreateTask	Creates a task and returns its details.
GetTask	Retrieves a task by its ID.
Current Limitations
Tasks are stored in memory and are lost when the server stops.
The server uses insecure transport credentials for local development.
There is no authentication or authorization.
The server does not yet use a persistent database.
Learning Goals

This project is intended to help beginners understand:

How gRPC client-server communication works.
How to define services and messages using Protocol Buffers.
How to generate Go code from .proto files.
How to implement RPC handlers in Go.
How to handle gRPC status codes and errors.

This project is for learning and experimentation.