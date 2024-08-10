module github.com/illenko/go-grpc-client

go 1.22.3

replace github.com/illenko/go-grpc-common => ../common

require (
	github.com/google/uuid v1.6.0
	github.com/illenko/go-grpc-common v0.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.47
	google.golang.org/grpc v1.65.0
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
