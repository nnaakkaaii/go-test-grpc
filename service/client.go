package service

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"test-grpc/pb"
	"test-grpc/pkg"
	"time"
)

func PlayGame(scanner *bufio.Scanner) {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// gRPCのクライアントの生成
	client := pb.NewRockPaperScissorsServiceClient(conn)

	// gRPCのリクエストの生成
	scanner.Scan()
	in := scanner.Text()

	switch in {
	case "1", "2", "3":
	default:
		fmt.Println("Invalid command.")
		return
	}

	handShapes, _ := strconv.Atoi(in)
	playRequest := pb.PlayRequest{HandShapes: pkg.EncodeHandShapes(int32(handShapes))}

	// コンテキストの生成
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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

	// gRPCのクライアントを生成
	client := pb.NewRockPaperScissorsServiceClient(conn)

	// gRPCのリクエストを生成
	reportRequest := pb.ReportRequest{}

	// コンテキストを生成
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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

func NotifyMessages(scanner *bufio.Scanner) {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// gRPCのクライアントの生成
	client := pb.NewRockPaperScissorsServiceClient(conn)

	// gRPCのリクエストの生成
	scanner.Scan()
	in := scanner.Text()

	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal("Parsing failed.")
		return
	}

	req := &pb.NotifyRequest{Num: int32(num)}

	// コンテキストの生成
	ctx := context.Background()

	// gRPCサーバーのNotifyMessagesメソッドを呼び出し
	stream, err := client.NotifyMessages(ctx, req)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}
	fmt.Println("***********************************")
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Stream failed.")
			return
		}
		fmt.Printf("Notification: %s \n", reply.GetMessage())
	}
	fmt.Println("***********************************")
	fmt.Println()
}

func SumValues(scanner *bufio.Scanner) {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// gRPCのクライアントの生成
	client := pb.NewRockPaperScissorsServiceClient(conn)

	// gRPCのリクエストの生成 & メソッドを呼び出し
	stream, err := client.SumValues(context.Background())
	if err != nil {
		log.Fatal("Stream failed.")
		return
	}
	fmt.Println("***********************************")
	for {
		scanner.Scan()
		in := scanner.Text()
		num, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal("Invalid values.")
			return
		}
		if num == 0 {
			break
		}
		fmt.Printf("received value: %d \n", num)
		if err := stream.Send(&pb.SumRequest{Value: int32(num)}); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Sending stream failed.")
			return
		}
		time.Sleep(1 * time.Second)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Closing stream failed.")
	}
	fmt.Printf("Results: %v \n", reply)
	fmt.Println("***********************************")
	fmt.Println()
}

func request(scanner *bufio.Scanner, stream pb.RockPaperScissorsService_ChatMessagesClient) error {
	for {
		scanner.Scan()
		in := scanner.Text()
		if in == "" {
			return nil
		}
		if err := stream.Send(&pb.ChatRequest{Message: in}); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Sending stream failed.")
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func receive(stream pb.RockPaperScissorsService_ChatMessagesClient) error {
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatal("Recv stream failed.")
			}
			fmt.Printf("From server: %s \n", in.Message)
		}
	}()
	<-waitc
	return nil
}

func ChatMessages(scanner *bufio.Scanner) {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// gRPCのクライアントの生成
	client := pb.NewRockPaperScissorsServiceClient(conn)

	// gRPCのリクエストの生成 & メソッドを呼び出し
	stream, err := client.ChatMessages(context.Background())
	if err != nil {
		log.Fatal("Stream failed.")
		return
	}
	fmt.Println("***********************************")
	if err := request(scanner, stream); err != nil {
		log.Fatal("Request failed.")
		return
	}
	if err := receive(stream); err != nil {
		log.Fatal("Response failed.")
		return
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatal("Close stream failed.")
		return
	}
	return
}
