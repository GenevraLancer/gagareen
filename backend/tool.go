// +build tools

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"                                   //simple configuration protoc, protoc-gen-* source and out dir
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway" //go module APIv2, generate reverse proxy to JSON
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"    //swagger
	_ "github.com/rakyll/statik"                                          //need to generate static assets for OpenAPI UI
	_ "golang.org/x/tools/gopls"                                          //go language server to orginise work with go modules within IDE
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"                     //grpc with go module APIv2
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"                      //go module APIv2
)
