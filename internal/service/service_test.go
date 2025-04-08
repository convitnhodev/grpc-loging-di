package service

// const bufSize = 1024 * 1024

// func TestServiceRegistration(t *testing.T) {
// 	// Create a buffer for our in-memory listener
// 	lis := bufconn.Listen(bufSize)

// 	// Mock config
// 	mockConfig := &config.Config{
// 		Name: "test-service",
// 		GrpcConfig: &grpc.Config{
// 			Port:              9999, // Not actually used since we're using bufconn
// 			EnableHealthCheck: true,
// 		},
// 	}

// 	// Create test context
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	logger, _ := zap.NewDevelopment()

// 	// Create the service with the mock config
// 	svc := NewService(ctx, mockConfig, logger)

// 	// Create a new gRPC server
// 	srv := gogrpc.NewServer()

// 	// Register the account service using the service's AccountService
// 	account.RegisterAccountServiceServer(srv, svc.AccountService)

// 	// Serve in goroutine
// 	go func() {
// 		if err := srv.Serve(lis); err != nil {
// 			t.Errorf("Server exited with error: %v", err)
// 		}
// 	}()

// 	// Create a client connection
// 	conn, err := gogrpc.DialContext(ctx, "bufnet", gogrpc.WithContextDialer(
// 		func(context.Context, string) (net.Conn, error) {
// 			return lis.Dial()
// 		}),
// 		gogrpc.WithTransportCredentials(insecure.NewCredentials()),
// 		gogrpc.WithBlock(),
// 	)

// 	assert.NoError(t, err, "Failed to dial bufnet")
// 	defer conn.Close()

// 	// Create client
// 	client := account.NewAccountServiceClient(conn)

// 	// Test the GetAccount method
// 	req := &account.GetAccountRequest{
// 		Id: "test-id",
// 	}

// 	resp, err := client.GetAccount(ctx, req)

// 	// Verify no error and response is as expected
// 	assert.NoError(t, err, "GetAccount failed")
// 	assert.NotNil(t, resp, "Response should not be nil")
// 	assert.Equal(t, int32(0), resp.Code, "Expected code 0")
// 	assert.Equal(t, "success", resp.Message, "Expected 'success' message")
// 	assert.Equal(t, "test-id", resp.Data.Id, "ID should match request")
// 	assert.Equal(t, "hung dep trai", resp.Data.Name, "Name should match expected value")
// }

// // TestServiceStart tests the service Start method
// func TestServiceStart(t *testing.T) {
// 	// Create a test context
// 	ctx := context.Background()

// 	// Create mock config
// 	mockConfig := &config.Config{
// 		Name: "test-service",
// 		GrpcConfig: &grpc.Config{
// 			Port:              0, // Use any available port
// 			EnableHealthCheck: true,
// 		},
// 	}

// 	logger, _ := zap.NewDevelopment()

// 	// Create service
// 	svc := NewService(ctx, mockConfig, logger)

// 	// Start service in goroutine
// 	errCh := make(chan error, 1)
// 	go func() {
// 		errCh <- svc.Start(ctx)
// 	}()

// 	// Give it time to start
// 	time.Sleep(100 * time.Millisecond)

// 	// Make sure no error occurred
// 	select {
// 	case err := <-errCh:
// 		assert.NoError(t, err, "Service Start returned error")
// 	default:
// 		// No error yet, service is still running (which is good)
// 	}

// 	// Stop the service
// 	err := svc.Stop(ctx)
// 	assert.NoError(t, err, "Service Stop returned error")
// }
