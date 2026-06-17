package main

import (
	"fmt"
	"net"
	"saas-monitoramento/backend/internal/alert"

	pb "saas-monitoramento/backend/gen"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("falha ao iniciar o servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	alertServer := &alert.AlertServer{}
	reflection.Register(grpcServer)
	pb.RegisterAlertServiceServer(grpcServer, alertServer)

	fmt.Println("Servidor gRPC rodando na porta 50051...")
	if err = grpcServer.Serve(lis); err != nil {
		fmt.Printf("falha ao servir: %v", err)
	}
}
