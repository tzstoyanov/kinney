include ../Makefile

GOOGLEAPIS=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

# Protocol Buffer Go API
api.pb.go:
	protoc -I. \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/${GOOGLEAPIS} \
		--go_out=plugins=grpc:. \
		--plugin=${GOBIN}/protoc-gen-go \
		api.proto

# HTTP REST to gRPC adapter
api.pb.gw.go:
	protoc -I. \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/${GOOGLEAPIS} \
		--grpc-gateway_out=logtostderr=true:. \
		--plugin=${GOBIN}/protoc-gen-grpc-gateway \
		api.proto

api_pb2.py:
	protoc -I. \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/${GOOGLEAPIS} \
		--python_out=. \
		api.proto

api_pb2_grpc.py:
	python -m grpc_tools.protoc -I. \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/${GOOGLEAPIS} \
		--grpc_python_out=. \
		api.proto

py: api_pb2.py api_pb2_grpc.py

proxy: api.pb.go api.pb.gw.go
	go install ${KINNEY}/orchestrator/proxy
	@echo Binary: ${GOBIN}/proxy

clean:
	rm -f *.pb*.go *_pb2*.py

.PHONY: env proxy clean py
