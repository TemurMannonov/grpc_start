package main

import (
	"context"
	"fmt"
	arifmeticService "grpc_start/genproto/arifmetic_service"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
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
