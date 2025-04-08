package main

import (
	"context"
	"fmt"
	"grpc/account/app"
	"grpc/account/config"
	account "grpc/spec/generated/account"
	"log"
	"time"

	"github.com/convitnhodev/common/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load config
	cfg := config.Load()
	if cfg == nil {
		log.Fatal("Failed to load config")
	}

	// Initialize application using Wire
	service, err := app.InitializeApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Give the service time to start
	time.Sleep(2 * time.Second)

	// Run test connection
	if err := testGRPCConnection(service.Logger()); err != nil {
		log.Fatalf("gRPC connection test failed: %v", err)
	}

	// Keep the main goroutine running
	select {}
}

func testGRPCConnection(logger logging.Logger) error {
	logger.Info("Connecting to gRPC server...")

	// Use WithBlock() to wait for connection
	conn, err := grpc.Dial(
		"localhost:9090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	logger.Info("Connection established")

	logger.Info("Creating client...")
	client := account.NewAccountServiceClient(conn)
	logger.Info("Client created")

	req := &account.GetAccountRequest{
		Id: "test-user-id",
	}
	logger.Info("Sending request", zap.String("request", fmt.Sprintf("%+v", req)))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetAccount(ctx, req)
	if err != nil {
		return fmt.Errorf("service call failed: %v", err)
	}

	logger.Info("Service response received",
		zap.Int32("code", resp.Code),
		zap.String("message", resp.Message),
		zap.Any("data", resp.Data),
	)

	// print context
	fmt.Println("Context:", ctx)

	logger.Info("gRPC connection test successful!")
	return nil
}
