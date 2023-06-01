# react-grpc-go
-Make proto file
-Generate protobuf files
  buf mod init
  buf generate
 ##Cach generate truc tiep bang protoc
  -generate G0
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/route_guide.proto
  -generate Typescript:
   protoc -I=. ./greet.proto --js_out=import_style=commonjs:./src/output --grpc-web_out=import_style=typescript,mode=grpcwebtext:./src/output
   
##import
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

