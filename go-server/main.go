package main

import (
	"context"
	"log"
	"net"

	pb "go-server/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserProfile(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if req.GetUserId() == 1 {
		return &pb.UserResponse{
			Name:  "Fajri Ayu",
			Email: "Ayu@universitas.ac.id",
			Phone: "081234567890",
		}, nil
	}

	return &pb.UserResponse{
		Name:  "User Tidak Ditemukan",
		Email: "",
		Phone: "",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Server gRPC berjalan di port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}