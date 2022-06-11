# go-test-grpc

## usage

### コードの自動生成 (gRPC-Gatewayなし)

```bash
$ protoc --go_out=. \
  --go-grpc_out=require_unimplemented_servers=false:. \
  ./proto/rock-paper-scissors.proto
```

### コードの自動生成 (gRPC-Gatewayあり)

gRPC stubの生成

```bash
$ protoc -I/usr/local/include -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
  -I/usr/local/opt/protobuf/include \
  --go_out=. \
  --go-grpc_out=require_unimplemented_servers=false:. \
  ./proto/rock-paper-scissors.proto
```

gRPC Gatewayの生成

```bash
$ protoc -I/usr/local/include -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
  -I/usr/local/opt/protobuf/include \
  --grpc-gateway_out=logtostderr=true:. \
  ./proto/rock-paper-scissors.proto
```

OpenAPI Schemaの生成

```bash
$ protoc -I/usr/local/include -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
  -I/usr/local/opt/protobuf/include \
  --swagger_out=allow_merge=true,merge_file_name=./spec:. \
  ./proto/rock-paper-scissors.proto
```

※ protoc-gen-grpc-gateway がnot foundになる場合、pkg下のprotoc-gen-grpc-gatewayに入ったあと、 `go install` を実行

### cliでの動作確認

gRPCの動作確認

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
$ grpc_cli call localhost:50051 game.RockPaperScissorsService.NotifyMessages 'num: 5'
connecting to localhost:50051
message: "0"
message: "1"
message: "2"
message: "3"
message: "4"
Rpc succeeded with OK status
$ grpc_cli call localhost:50051 game.RockPaperScissorsService.SumValues
reading streaming request message from stdin...
value: 10

Request sent.
value: 11

Request sent.
```

Rest APIの確認

```bash
$ curl -XGET "localhost:50052/v1/results"
...
$ curl -XPOST -d '{"handShapes": 1}' "localhost:50052/v1/game/play"
...
$ curl -XGET "localhost:50052/v1/messages/notify?num=5"
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
