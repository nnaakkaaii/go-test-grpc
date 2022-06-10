package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"test-grpc/pb"
	"test-grpc/pkg"
	"time"
)

func PlayGame(handShapes int32) {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// gRPCのクライアントとリクエストの生成
	client := pb.NewRockPaperScissorsServiceClient(conn)
	playRequest := pb.PlayRequest{HandShapes: pkg.EncodeHandShapes(handShapes)}

	// gRPCサーバーのPlayGameメソッドを呼び出し
	reply, err := client.PlayGame(ctx, &playRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	// レスポンスを標準出力に表示
	marchResult := reply.GetMatchResult()
	fmt.Println("***********************************")
	fmt.Printf("Your hand shapes: %s \n", marchResult.YourHandShapes.String())
	fmt.Printf("Opponent hand shapes: %s \n", marchResult.OpponentHandShapes.String())
	fmt.Printf("Result: %s \n", marchResult.Result.String())
	fmt.Println("***********************************")
	fmt.Println()
}

func ReportMatchResults() {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// gRPCのクライアントとリクエストを生成
	client := pb.NewRockPaperScissorsServiceClient(conn)
	reportRequest := pb.ReportRequest{}

	// gRPCサーバーのReportMatchResultsメソッドを呼び出し
	reply, err := client.ReportMatchResults(ctx, &reportRequest)
	if err != nil {
		log.Fatal("Request Failed.")
		return
	}

	// レスポンスを標準出力に表示
	report := reply.GetReport()
	if len(report.MatchResults) == 0 {
		fmt.Println("***********************************")
		fmt.Println("There are no match results.")
		fmt.Println("***********************************")
		fmt.Println()
		return
	}

	fmt.Println("***********************************")
	for k, v := range report.MatchResults {
		fmt.Println(k + 1)
		fmt.Printf("Your hand shapes: %s \n", v.YourHandShapes.String())
		fmt.Printf("Opponent hand shapes: %s \n", v.OpponentHandShapes.String())
		fmt.Printf("Result: %s \n", v.Result.String())
		fmt.Printf("Datetime of match: %s \n", v.CreateTime.AsTime().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format(time.ANSIC))
		fmt.Println()
	}

	fmt.Printf("Number of games: %d \n", reply.GetReport().NumberOfGames)
	fmt.Printf("Number of wins: %d \n", reply.GetReport().NumberOfWins)
	fmt.Println("***********************************")
	fmt.Println()
}
