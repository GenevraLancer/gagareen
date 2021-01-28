# generate заменена функцией givemeproto c buf плагином
# generate:
	# Generate go, gRPC-Gateway, swagger
	# 
	# --proto_path source folder
	#
	# -I declares import folders
	#
	# --protoc-gen-go_out generates GoGo Protobuf output with gRPC plugin enabled.
	# --protoc-gen-grpc-gateway_out generates gRPC-Gateway output.
	# --protoc-gen-openapiv2_out generates swagger (an OpenAPI 2.0 specification) for our gRPC-Gateway endpoints.
	#
	# ./gen is the output directory.
	#
	# ./protos is the source location
#	protoc \
		--go_out=../server/gen/api \
		--go_opt=paths=source_relative \
		--go-grpc_out=../server/gen/api \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=paths=source_relative,logtostderr=true:. \
		--openapiv2_out=../server/gen/swagger \
		--js_out=import_style=commonjs,binary:../client/src/ \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:../client/src/ \
		./protos/*.proto \


	# Generate static assets for OpenAPI UI
	#statik -m -f -src third_party/OpenAPI/
givemeproto:
	buf generate --path ./protos/*.proto
	# Generate static assets for OpenAPI UI
	statik -m -f -src third_party/OpenAPI/

install:
	go get \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    	google.golang.org/protobuf/cmd/protoc-gen-go \
    	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/rakyll/statik \
		github.com/bufbuild/buf/cmd/buf \
		golang.org/x/tools/gopls

	# compile:
		# echo "Compiling for every OS and Platform"
		# GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
		# GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
		# GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
run-server:
	go run main.go

build-dev:
	go build ./...

build-prod: givemeproto
	go build -o /bin/gagarin

download:
  @echo Download go.mod dependencies
  @go mod download

install-tools: download
  @echo Installing tools from tools.go
  @cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %