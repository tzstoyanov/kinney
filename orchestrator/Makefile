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
	rm -f *_pb2*.py

.PHONY: clean py
