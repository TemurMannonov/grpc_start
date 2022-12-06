package main

import (
	"context"
	"fmt"
	"net"

	arifmeticService "github.com/TemurMannonov/grpc_start/genproto/arifmetic_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	arifmeticService.UnimplementedArifmeticServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	arifmeticService.RegisterArifmeticServiceServer(srv, &Server{})
	reflection.Register(srv)

	fmt.Println("server started")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *Server) Add(ctx context.Context, request *arifmeticService.Request) (*arifmeticService.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &arifmeticService.Response{Result: result}, nil
}

func (s *Server) Multiply(ctx context.Context, request *arifmeticService.Request) (*arifmeticService.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &arifmeticService.Response{Result: result}, nil
}

func (s *Server) Substract(ctx context.Context, request *arifmeticService.Request) (*arifmeticService.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b

	return &arifmeticService.Response{Result: result}, nil
}

func (s *Server) Divide(ctx context.Context, request *arifmeticService.Request) (*arifmeticService.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a / b

	return &arifmeticService.Response{Result: result}, nil
}
