GOPATH_DIR := $(GOPATH)/src/github.com/open-telemetry/opentelemetry-proto
GENDIR := gen
GOPATH_GENDIR := $(GOPATH_DIR)/$(GENDIR)

# Find all .proto files.
PROTO_FILES := $(wildcard opentelemetry/proto/*/*/*.proto opentelemetry/proto/*/*/*/*.proto)

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

# Generate all implementations
.PHONY: gen-all
gen-all: gen-cpp gen-csharp gen-go gen-java gen-kotlin gen-objc gen-openapi gen-php gen-python gen-ruby

OTEL_DOCKER_PROTOBUF ?= otel/build-protobuf:0.9.0
PROTOC := docker run --rm -u ${shell id -u} -v${PWD}:${PWD} -w${PWD} ${OTEL_DOCKER_PROTOBUF} --proto_path=${PWD}
PROTO_INCLUDES := -I/usr/include/github.com/gogo/protobuf
PROTO_FLAGS := --experimental_allow_proto3_optional

PROTO_GEN_CPP_DIR ?= $(GENDIR)/cpp
PROTO_GEN_CSHARP_DIR ?= $(GENDIR)/csharp
PROTO_GEN_GO_DIR ?= $(GENDIR)/go
PROTO_GEN_JAVA_DIR ?= $(GENDIR)/java
PROTO_GEN_JS_DIR ?= $(GENDIR)/js
PROTO_GEN_KOTLIN_DIR ?= $(GENDIR)/kotlin
PROTO_GEN_OBJC_DIR ?= $(GENDIR)/objc
PROTO_GEN_OPENAPI_DIR ?= $(GENDIR)/openapi
PROTO_GEN_PHP_DIR ?= $(GENDIR)/php
PROTO_GEN_PYTHON_DIR ?= $(GENDIR)/python
PROTO_GEN_RUBY_DIR ?= $(GENDIR)/ruby

# Docker pull image.
.PHONY: docker-pull
docker-pull:
	docker pull $(OTEL_DOCKER_PROTOBUF)

# Generate gRPC/Protobuf implementation for C++.
.PHONY: gen-cpp
gen-cpp:
	rm -rf ./$(PROTO_GEN_CPP_DIR)
	mkdir -p ./$(PROTO_GEN_CPP_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS) --cpp_out=./$(PROTO_GEN_CPP_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --cpp_out=./$(PROTO_GEN_CPP_DIR) --grpc-cpp_out=./$(PROTO_GEN_CPP_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --cpp_out=./$(PROTO_GEN_CPP_DIR) --grpc-cpp_out=./$(PROTO_GEN_CPP_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --cpp_out=./$(PROTO_GEN_CPP_DIR) --grpc-cpp_out=./$(PROTO_GEN_CPP_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf implementation for C#.
.PHONY: gen-csharp
gen-csharp:
	rm -rf ./$(PROTO_GEN_CSHARP_DIR)
	mkdir -p ./$(PROTO_GEN_CSHARP_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS) --csharp_out=./$(PROTO_GEN_CSHARP_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --csharp_out=./$(PROTO_GEN_CSHARP_DIR) --grpc-csharp_out=./$(PROTO_GEN_CSHARP_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --csharp_out=./$(PROTO_GEN_CSHARP_DIR) --grpc-csharp_out=./$(PROTO_GEN_CSHARP_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --csharp_out=./$(PROTO_GEN_CSHARP_DIR) --grpc-csharp_out=./$(PROTO_GEN_CSHARP_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf implementation for Go.
.PHONY: gen-go
gen-go:
	rm -rf ./$(PROTO_GEN_GO_DIR)
	mkdir -p ./$(PROTO_GEN_GO_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command,$(PROTOC) $(PROTO_FLAGS) $(PROTO_INCLUDES) --gogo_out=plugins=grpc:./$(PROTO_GEN_GO_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=opentelemetry/proto/collector/trace/v1/trace_service_http.yaml:./$(PROTO_GEN_GO_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=opentelemetry/proto/collector/metrics/v1/metrics_service_http.yaml:./$(PROTO_GEN_GO_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --grpc-gateway_out=logtostderr=true,grpc_api_configuration=opentelemetry/proto/collector/logs/v1/logs_service_http.yaml:./$(PROTO_GEN_GO_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf implementation for Java.
.PHONY: gen-java
gen-java:
	rm -rf ./$(PROTO_GEN_JAVA_DIR)
	mkdir -p ./$(PROTO_GEN_JAVA_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS) --java_out=./$(PROTO_GEN_JAVA_DIR) $(file)))

# Generate gRPC/Protobuf implementation for Kotlin.
.PHONY: gen-kotlin
gen-kotlin: gen-java
	rm -rf ./$(PROTO_GEN_KOTLIN_DIR)
	mkdir -p ./$(PROTO_GEN_KOTLIN_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) --kotlin_out=./$(PROTO_GEN_KOTLIN_DIR) $(file)))


# Generate gRPC/Protobuf implementation for JavaScript.
.PHONY: gen-js
gen-js:
	rm -rf ./$(PROTO_GEN_JS_DIR)
	mkdir -p ./$(PROTO_GEN_JS_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS)  --js_out=import_style=commonjs:./$(PROTO_GEN_JS_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --js_out=import_style=commonjs:./$(PROTO_GEN_JS_DIR) --grpc-web_out=import_style=commonjs,mode=grpcweb:./$(PROTO_GEN_JS_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --js_out=import_style=commonjs:./$(PROTO_GEN_JS_DIR) --grpc-web_out=import_style=commonjs,mode=grpcweb:./$(PROTO_GEN_JS_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --js_out=import_style=commonjs:./$(PROTO_GEN_JS_DIR) --grpc-web_out=import_style=commonjs,mode=grpcweb:./$(PROTO_GEN_JS_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf implementation for Objective-C.
.PHONY: gen-objc
gen-objc:
	rm -rf ./$(PROTO_GEN_OBJC_DIR)
	mkdir -p ./$(PROTO_GEN_OBJC_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS) --objc_out=./$(PROTO_GEN_OBJC_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --objc_out=./$(PROTO_GEN_OBJC_DIR) --grpc-objc_out=./$(PROTO_GEN_OBJC_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --objc_out=./$(PROTO_GEN_OBJC_DIR) --grpc-objc_out=./$(PROTO_GEN_OBJC_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --objc_out=./$(PROTO_GEN_OBJC_DIR) --grpc-objc_out=./$(PROTO_GEN_OBJC_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf for openapi v2 (swagger)
.PHONY: gen-openapi
gen-openapi:
	mkdir -p $(PROTO_GEN_OPENAPI_DIR)
	$(PROTOC) $(PROTO_FLAGS) --openapiv2_out=logtostderr=true,grpc_api_configuration=opentelemetry/proto/collector/trace/v1/trace_service_http.yaml:$(PROTO_GEN_OPENAPI_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --openapiv2_out=logtostderr=true,grpc_api_configuration=opentelemetry/proto/collector/metrics/v1/metrics_service_http.yaml:$(PROTO_GEN_OPENAPI_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --openapiv2_out=logtostderr=true,grpc_api_configuration=opentelemetry/proto/collector/logs/v1/logs_service_http.yaml:$(PROTO_GEN_OPENAPI_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf implementation for PhP.
.PHONY: gen-php
gen-php:
	rm -rf ./$(PROTO_GEN_PHP_DIR)
	mkdir -p ./$(PROTO_GEN_PHP_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS) --php_out=./$(PROTO_GEN_PHP_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --php_out=./$(PROTO_GEN_PHP_DIR) --grpc-php_out=./$(PROTO_GEN_PHP_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --php_out=./$(PROTO_GEN_PHP_DIR) --grpc-php_out=./$(PROTO_GEN_PHP_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --php_out=./$(PROTO_GEN_PHP_DIR) --grpc-php_out=./$(PROTO_GEN_PHP_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf implementation for Python.
.PHONY: gen-python
gen-python:
	rm -rf ./$(PROTO_GEN_PYTHON_DIR)
	mkdir -p ./$(PROTO_GEN_PYTHON_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS)  --python_out=./$(PROTO_GEN_PYTHON_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --python_out=./$(PROTO_GEN_PYTHON_DIR) --grpc-python_out=./$(PROTO_GEN_PYTHON_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --python_out=./$(PROTO_GEN_PYTHON_DIR) --grpc-python_out=./$(PROTO_GEN_PYTHON_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --python_out=./$(PROTO_GEN_PYTHON_DIR) --grpc-python_out=./$(PROTO_GEN_PYTHON_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto

# Generate gRPC/Protobuf implementation for Ruby.
.PHONY: gen-ruby
gen-ruby:
	rm -rf ./$(PROTO_GEN_RUBY_DIR)
	mkdir -p ./$(PROTO_GEN_RUBY_DIR)
	$(foreach file,$(PROTO_FILES),$(call exec-command, $(PROTOC) $(PROTO_FLAGS) --ruby_out=./$(PROTO_GEN_RUBY_DIR) $(file)))
	$(PROTOC) $(PROTO_FLAGS) --ruby_out=./$(PROTO_GEN_RUBY_DIR) --grpc-ruby_out=./$(PROTO_GEN_RUBY_DIR) opentelemetry/proto/collector/trace/v1/trace_service.proto
	$(PROTOC) $(PROTO_FLAGS) --ruby_out=./$(PROTO_GEN_RUBY_DIR) --grpc-ruby_out=./$(PROTO_GEN_RUBY_DIR) opentelemetry/proto/collector/metrics/v1/metrics_service.proto
	$(PROTOC) $(PROTO_FLAGS) --ruby_out=./$(PROTO_GEN_RUBY_DIR) --grpc-ruby_out=./$(PROTO_GEN_RUBY_DIR) opentelemetry/proto/collector/logs/v1/logs_service.proto
