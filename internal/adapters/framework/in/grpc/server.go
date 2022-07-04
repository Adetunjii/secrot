package grpc

import (
	"fmt"
	"github.com/Adetunjii/secrot/internal/adapters/framework/in/grpc/pb"
	ports "github.com/Adetunjii/secrot/internal/ports/in"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpcAdapter *Adapter) Run() {
	var err error

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("port:: %s", port)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	userServiceServer := grpcAdapter
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc server over port %s: %v", port, err)
	} else {
		log.Printf("server is running on port: %s", port)
	}

}
