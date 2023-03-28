package main

import (
	"database/sql"
	"github.com/PGabrielDev/grpc-go/interal/database"
	"github.com/PGabrielDev/grpc-go/internal/pb"
	"github.com/PGabrielDev/grpc-go/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	db, err := sql.Open("sqlite3", "./grpc.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(categoryDB)
	grpcService := grpc.NewServer()
	reflection.Register(grpcService)
	pb.RegisterCategoryServiceServer(grpcService, categoryService)
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err = grpcService.Serve(listen); err != nil {
		panic(err)
	}
}
