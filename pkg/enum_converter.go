package pkg

import "test-grpc/pb"

func EncodeHandShapes(n int32) pb.HandShapes {
	switch n {
	case 1:
		return pb.HandShapes_ROCK
	case 2:
		return pb.HandShapes_PAPER
	case 3:
		return pb.HandShapes_SCISSORS
	default:
		return pb.HandShapes_HAND_SHAPES_UNKNOWN
	}
}
