module github.com/tmpgat/server

go 1.19

require (
	github.com/tmpgat/pb v0.0.0
	github.com/bufbuild/connect-go v1.4.1
	github.com/rs/cors v1.8.3
	golang.org/x/net v0.5.0
	google.golang.org/protobuf v1.28.1
)

require golang.org/x/text v0.6.0 // indirect

replace (
	github.com/tmpgat/pb v0.0.0 => ./pb/gen/go
)
