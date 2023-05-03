@echo off
pushd %~dp0\..\
echo Generating proto buffers...
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/proto/hello/hello.proto pkg/proto/jwt/jwt.proto pkg/proto/signup/signup.proto
popd
PAUSE