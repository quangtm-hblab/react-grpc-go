# react-grpc-go
-Make proto file
-Generate protobuf files
  buf mod init
  buf generate
-Import generated file to react-client, go-server
-Imlement client, server, proto services in server
-server listen at port 50051(HTTP2.0), client call request at port 8080(HTTP 1.1)
-create envoy proxy server listen at port 8080, call at port 50051
-run go-server
  go mod tidy
  go run main.go
-reactclient
  npm install
  npm start
-envoy proxy server
  envoy -c ./path-to-envoy.yaml

