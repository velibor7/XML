module github.com/velibor7/XML/api_gateway

go 1.18

replace github.com/velibor7/XML/common => ../common

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/sirupsen/logrus v1.8.1
	github.com/velibor7/XML/common v1.0.0
	google.golang.org/grpc v1.46.2
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.12.2
	gitlab.com/msvechla/mux-prometheus v0.0.2
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
