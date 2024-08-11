module github.com/illenko/go-grpc-common

go 1.22.0

require (
    google.golang.org/grpc v1.65.0
    google.golang.org/protobuf v1.34.2
    github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
    google.golang.org/genproto v0.0.0-20240528184218-531527333157
)

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20240528184218-531527333157