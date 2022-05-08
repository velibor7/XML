module github.com/velibor7/XML/connection_service

go 1.18

replace github.com/velibor7/XML/common => ../common

require (
	github.com/neo4j/neo4j-go-driver/v4 v4.4.2
	github.com/velibor7/XML/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.46.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
