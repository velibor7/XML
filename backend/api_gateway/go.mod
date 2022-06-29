module github.com/velibor7/XML/api_gateway

go 1.18

replace github.com/velibor7/XML/common => ../common

require (
	github.com/gorilla/handlers v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/velibor7/XML/common v1.0.0
	google.golang.org/grpc v1.46.2
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
