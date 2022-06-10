package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"test-grpc/pb"
	"test-grpc/service"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーの生成
	server := grpc.NewServer()
	// 自動生成された関数に、サーバーと実際に処理を行うメソッドを実装したハンドラを設定
	pb.RegisterRockPaperScissorsServiceServer(server, service.NewRockPaperScissorsService())

	// サーバーリフレクションを有効にすることで、シリアライズせずにgrpc_cli上での動作確認ができる
	reflection.Register(server)
	// サーバーを起動
	server.Serve(listenPort)
}
