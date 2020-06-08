# This Makefile is assumed to be in the root of the repository, so its location
# is used to derive the directory that is the repository root.
#
# https://www.gnu.org/software/make/manual/html_node/Special-Variables.html#index-MAKEFILE_005fLIST-_0028list-of-parsed-makefiles_0029
# https://www.gnu.org/software/make/manual/html_node/Text-Functions.html#index-firstword
# https://www.gnu.org/software/make/manual/html_node/File-Name-Functions.html#index-dir
REPO_ROOT := $(dir $(firstword $(MAKEFILE_LIST)))

################################################################################
## Go
################################################################################

GO_MODULE_NAME := github.com/CamusEnergy/kinney

################################################################################
## Python
################################################################################

# Initializes the Python virtual environment for development, or installs new
# dependencies into the existing venv, as appropriate.
pipenv:
	pipenv install --dev
	# Install a "*.pth" file in the virtual environment's site-packages
	# directory in order to support importing code from the local repository
	# without having to bundle and install the Python package.
	pipenv run python util/install_pth.py \
		--package_name="kinney" \
		--package_root="$(REPO_ROOT)"
.PHONY: pipenv

################################################################################
## Protocol Buffers
################################################################################

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
	protoc \
		--go_out="module=${GO_MODULE_NAME}:." \
		--plugin="./bin/protoc-gen-go" \
		"$<"

# Generate the corresponding Go gRPC source for the services in each Protocol
# Buffer descriptor file.  If the input file contains no service definitions,
# then no output file will be created.
#
# Note that this pattern is a subset of the `%.pb.go` rule: make "will choose
# the rule with the shortest stem (that is, the pattern that matches most
# specifically)."
# https://www.gnu.org/software/make/manual/html_node/Pattern-Match.html#Pattern-Match
%_grpc.pb.go: %.proto bin/protoc-gen-go-grpc
	protoc \
		--go-grpc_out="module=${GO_MODULE_NAME}:." \
		--plugin="./bin/protoc-gen-go-grpc" \
		"$<"

# Generate the corresponding Python source for each Protocol Buffer descriptor
# file.
%_pb2.py: %.proto
	protoc --python_out="." "$<"

# Generate the corresponding Python gRPC source for the services in each
# Protocol Buffer descriptor file.
#
# The Python gRPC generator has to run inside of the Python venv environment, as
# it is installed by the `grpcio-tools` Python package:
# https://grpc.io/docs/languages/python/basics/#generating-client-and-server-code
#
# Note that, contrary to the Go gRPC generator, this *will* create an output
# file (that is effectively empty) if the input file contains no service
# definitions.
%_pb2_grpc.py: %.proto pipenv
	pipenv run python -m grpc_tools.protoc \
		--proto_path="." \
		--grpc_python_out="." \
		"$<"

# Helper rule to regenerate all Protocol Buffer sources at once.
protos: orchestrator/api.pb.go orchestrator/api_grpc.pb.go
protos: orchestrator/api_pb2.py orchestrator/api_pb2_grpc.py
.PHONY: protos

################################################################################
## Cleanup
################################################################################

clean:
	# Remove the directory containing locally installed tools.
	-rm -r ./bin/
	# Delete the Python virtual environment.
	-pipenv --rm
.PHONY: clean
