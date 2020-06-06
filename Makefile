GO_MODULE="github.com/CamusEnergy/kinney"

# Install the Go protoc plugin to `./bin/protoc-gen-go`.
bin/protoc-gen-go:
	go build -o "$@" "google.golang.org/protobuf/cmd/protoc-gen-go"

# Install the Go gRPC protoc plugin to `./bin/protoc-gen-go-grpc`.
bin/protoc-gen-go-grpc:
	go build -o "$@" "google.golang.org/grpc/cmd/protoc-gen-go-grpc"

# Generate the corresponding Go source for each Protocol Buffer descriptor file.
#
# Note that the "module" parameter to the plugin is an undocumented feature that
# strips the module name prefix off of the output filenames.  Without this, all
# of the output files would end up in tree rooted at "github.com":
# https://github.com/protocolbuffers/protobuf-go/blob/69839c7/compiler/protogen/protogen.go#L430
%.pb.go: %.proto bin/protoc-gen-go
	protoc --go_out="module=${GO_MODULE}:." --plugin="./bin/protoc-gen-go" "$<"

# Generate the corresponding Go gRPC source for the services in each Protocol
# Buffer descriptor file.
#
# Note that this pattern is a subset of the `%.pb.go` rule: make "will choose
# the rule with the shortest stem (that is, the pattern that matches most
# specifically)."
# https://www.gnu.org/software/make/manual/html_node/Pattern-Match.html#Pattern-Match
%_grpc.pb.go: %.proto bin/protoc-gen-go-grpc
	protoc --go-grpc_out="module=${GO_MODULE}:." --plugin="./bin/protoc-gen-go-grpc" "$<"

clean:
	rm -r ./bin/
	find . -name '*.pb.go' -delete
.PHONY: clean

pipenv:
	pip install pipenv
	pipenv install
	@if [ -z "${PIPENV_ACTIVE}" ]; \
	then \
		echo ====================================; \
		echo WARNING: Currently outside of pipenv; \
		echo Run to enter: \`pipenv shell\`; \
		echo ====================================; \
	fi
.PHONY: pipenv
