package pkg

import "test-grpc/pb"

func JudgeWinOrLose(you, opponent pb.HandShapes) pb.Result {
	switch true {
	case you == opponent:
		return pb.Result_DRAW
	case (you.Number()-opponent.Number()+3)%3 == 1:
		return pb.Result_WIN
	default:
		return pb.Result_LOSE
	}
}
