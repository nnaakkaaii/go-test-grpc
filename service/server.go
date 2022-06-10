package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"math/rand"
	"net"
	"test-grpc/pb"
	"test-grpc/pkg"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StartServer(serverPort int) error {
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	// gRPCサーバーの生成
	server := grpc.NewServer()
	// 自動生成された関数に、サーバーと実際に処理を行うメソッドを実装したハンドラを設定
	pb.RegisterRockPaperScissorsServiceServer(server, newRockPaperScissorsService())

	// サーバーリフレクションを有効にすることで、シリアライズせずにgrpc_cli上での動作確認ができる
	reflection.Register(server)
	// サーバーを起動
	return server.Serve(listenPort)
}

var _ pb.RockPaperScissorsServiceServer = (*RockPaperScissorsService)(nil)

// RockPaperScissorsService では、DBを使わずに対戦結果の履歴を表示できるように構造体にデータを保持する
type RockPaperScissorsService struct {
	numberOfGames int32
	numberOfWins  int32
	matchResults  []*pb.MatchResult
}

// NewRockPaperScissorsService は、RockPaperScissorsServicesを生成するコンストラクタ
func newRockPaperScissorsService() *RockPaperScissorsService {
	return &RockPaperScissorsService{
		numberOfGames: 0,
		numberOfWins:  0,
		matchResults:  make([]*pb.MatchResult, 0),
	}
}

func (s *RockPaperScissorsService) PlayGame(ctx context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	if req.HandShapes == pb.HandShapes_HAND_SHAPES_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	// ランダムに1=3の数値を生成して相手の手とし、HandShapesのenumに変換
	opponentHandShapes := pkg.EncodeHandShapes(int32(rand.Intn(3) + 1))

	// ジャンケンの勝敗を決定
	result := pkg.JudgeWinOrLose(req.HandShapes, opponentHandShapes)

	now := time.Now()
	// 対戦結果を生成
	matchResult := &pb.MatchResult{
		YourHandShapes:     req.HandShapes,
		OpponentHandShapes: opponentHandShapes,
		Result:             result,
		CreateTime: &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	// 試合数を1増やし、プレイヤーが勝利した場合は勝利数も1増やす
	s.numberOfGames++
	if result == pb.Result_WIN {
		s.numberOfWins++
	}
	s.matchResults = append(s.matchResults, matchResult)

	return &pb.PlayResponse{MatchResult: matchResult}, nil
}

func (s *RockPaperScissorsService) ReportMatchResults(ctx context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	return &pb.ReportResponse{Report: &pb.Report{
		NumberOfGames: s.numberOfGames,
		NumberOfWins:  s.numberOfWins,
		MatchResults:  s.matchResults,
	}}, nil
}
