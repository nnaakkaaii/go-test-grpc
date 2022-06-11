.PHONY: gen-stub
gen-stub:
	protoc -I/usr/local/include -I. \
      -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
      -I/usr/local/opt/protobuf/include \
      --go_out=. \
      --go-grpc_out=require_unimplemented_servers=false:. \
      ./proto/rock-paper-scissors.proto

.PHONY: gen-gw
gen-gw:
	protoc -I/usr/local/include -I. \
      -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
      -I/usr/local/opt/protobuf/include \
      --grpc-gateway_out=logtostderr=true:. \
      ./proto/rock-paper-scissors.proto

.PHONY: gen-spec
gen-spec:
	protoc -I/usr/local/include -I. \
      -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
      -I/usr/local/opt/protobuf/include \
      --swagger_out=allow_merge=true,merge_file_name=./spec:. \
      ./proto/rock-paper-scissors.proto