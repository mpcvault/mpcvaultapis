# MPCVault APIs

This repo contains public interface definitions of MPCVault APIs.

MPCVault APIs use Protocol Buffers version 3 (proto3) as the Interface Definition Language (IDL) to define the API interface and the structure of the payload messages.

You can access MPCVault APIs published in this repository through GRPC, which is a high-performance binary RPC protocol over HTTP/2. It offers many useful features, including request/response multiplex and full-duplex streaming.

## Generate gRPC Source Code and Client Libraries

To create gRPC bindings for MPCVault APIs within this repository, you need to install Protocol Buffers and gRPC on your local machine first.

For a seamless start with gRPC, we suggest referring to the official tutorials [here](https://grpc.io/docs/languages).

As an example, if you aim to compile GoLang bindings, you should first install Go, protoc, and Go-specific protocol compiler plugins as detailed [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites).

Afterwards, you can compile the GoLang bindings using the command given below:

```bash
# protoc -I . --go_out=plugins=grpc:./genproto --go_opt=paths=source_relative \
#    --grpc-gateway_out=./genproto --grpc-gateway_opt=paths=source_relative \
#    mpcvault/platform/v1/*.proto

# 1. Clone the forked connect-openapi repository (includes the no-schema-title patch)
git clone https://github.com/lucastq1019/protoc-gen-connect-openapi.git

# 2. Build and install the plugin binary
cd protoc-gen-connect-openapi
go build -o $(go env GOPATH)/bin/protoc-gen-connect-openapi .
cd ..

# 3. Generate the OpenAPI spec
protoc -I . \
  --connect-openapi_out=gen \
  --connect-openapi_opt=format=json \
  --connect-openapi_opt=path-prefix= \
  --connect-openapi_opt=no-schema-title \
  --connect-openapi_opt=base=base.yaml \
  mpcvault/platform/v1/api.proto

# 4. Copy the generated output to the Mintlify docs directory
cp gen/mpcvault/platform/v1/api.openapi.json mintlify-docs/openapi.json

```