# react-grpc-go
- create .proto file
## [Opt 1] Generate protobuf files by [buf](https://github.com/bufbuild/buf)
  - init buf project \
  `buf mod init`
  - create file buf.gen.yaml 
  - generate protobuf \
  `buf generate`
## [Opt 2] Generate protobuf files by protoc
  - Generate Go protobuf [require plugin protoc-gen-go]:\
   `protoc --go_out=/dir/out --go_opt=paths=source_relative --go-grpc_out=dir/out --go-grpc_opt=paths=source_relative path/to/file.proto`
  - Generate Typescript protobuf: [require plugin protoc-gen-grpc-web]\
   `protoc -I=. ./greet.proto --js_out=import_style=commonjs:./src/output --grpc-web_out=import_style=typescript,mode=grpcwebtext:./src/output`
## Import protobuf files to server/client
  - Import generated file to react-client, go-server
  - Imlement client, server, proto services in server
  - server listen at port 50051(HTTP2.0), client call request at port 8080(HTTP 1.1)
  - create envoy proxy server listen at port 8080, call at port 50051
-run go-server
  go mod tidy
  go run main.go
-reactclient
  npm install
  npm start
-envoy proxy server
  envoy -c ./path-to-envoy.yaml

