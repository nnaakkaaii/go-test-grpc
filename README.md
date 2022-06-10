# go-test-grpc

## usage

### コードの自動生成

```bash
$ protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/rock-paper-scissors.proto
```

### cliでの動作確認

```bash
$ go run ./cmd/api &
$ grpc_cli ls localhost:50051 game.RockPaperScissorsService
PlayGame
ReportMatchResults
$ grpc_cli call localhost:50051 game.RockPaperScissorsService.PlayGame 'handShapes: 1'
connecting to localhost:50051
matchResult {
  yourHandShapes: ROCK
  opponentHandShapes: SCISSORS
  result: WIN
  create_time {
    seconds: 1632749016
    nanos: 862928000
  }
}
Rpc succeeded with OK status
$ grpc_cli call localhost:50051 game.RockPaperScissorsService.ReportMatchResults ''
connecting to localhost:50051
report {
  numberOfGames: 1
  numberOfWins: 1
  matchResults {
    yourHandShapes: ROCK
    opponentHandShapes: SCISSORS
    result: WIN
    create_time {
      seconds: 1632749016
      nanos: 862928000
    }
  }
}
Rpc succeeded with OK status
```

### clientからの確認

```bash
$ go run ./cmd/api &
$ go run ./cmd/cli
start Rock-paper-scissors game.
1: play game
2: show match results
3: exit
```