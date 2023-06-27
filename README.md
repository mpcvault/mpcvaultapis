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
protoc --go_out=./genproto --go_opt=paths=source_relative \
    --go-grpc_out=./genproto --go-grpc_opt=paths=source_relative \
    mpcvault/platform/v1/*.proto
```