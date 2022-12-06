package main

import (
	"context"
	"fmt"

	arifmeticService "github.com/TemurMannonov/grpc_start/genproto/arifmetic_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := arifmeticService.NewArifmeticServiceClient(conn)

	resp, err := client.Add(context.Background(), &arifmeticService.Request{A: 1231, B: 345})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result: ", resp.Result)

	resp, err = client.Divide(context.Background(), &arifmeticService.Request{A: 1231, B: 345})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result: ", resp)
}
