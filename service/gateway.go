package service

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
	"test-grpc/pb"
)

func StartGateway(serverPort int, gwPort int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(":%d", serverPort)
	if err := pb.RegisterRockPaperScissorsServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", gwPort), mux)
}
