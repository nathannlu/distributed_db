package main


import (
	"fmt"
	"github.com/nathannlu/distributed_db/conn"
)

func main() {

	// Expose http port for r/w queries from clients
	conn.StartHttpServer();

	// Initialize GRPC ports
	conn.StartGrpcServer();


	// Connect GRPC to form raft group
	

	






	fmt.Println("Distributed database is running")
}
