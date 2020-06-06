include ../Makefile

# Protocol Buffer Go API
api.pb.go:
	protoc -I. \
		--go_out=plugins=grpc:. \
		--plugin=${GOBIN}/protoc-gen-go \
		api.proto

api_pb2.py:
	protoc -I. \
		--python_out=. \
		api.proto

api_pb2_grpc.py:
	python -m grpc_tools.protoc -I. \
		--grpc_python_out=. \
		api.proto

py: api_pb2.py api_pb2_grpc.py

clean:
	rm -f *.pb*.go *_pb2*.py

.PHONY: env clean py
