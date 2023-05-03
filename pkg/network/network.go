package network

import (
	"fmt"
	"log"
	"net"

	"github.com/xd-Abi/moxie/pkg/logging"
	"google.golang.org/grpc"
)

type MicroServiceServer struct {
	port           uint
	log            *logging.Log
	listener       *net.Listener
	InternalServer *grpc.Server
}

func NewMicroServiceServer(port uint, log *logging.Log) *MicroServiceServer {
	log.Info("Initializing server...")
	log.Info("Creating tcp listener on port %v...", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		log.Fatal("Failed to listen on port %v: %v", port, err)
	}

	return &MicroServiceServer{
		log:            log,
		listener:       &listener,
		port:           port,
		InternalServer: grpc.NewServer(),
	}
}

func (s *MicroServiceServer) Start() {
	s.log.Info("Starting service server 🚀")
	if err := s.InternalServer.Serve(*s.listener); err != nil {
		log.Fatal("Failed to serve grpc server")
	}
}
