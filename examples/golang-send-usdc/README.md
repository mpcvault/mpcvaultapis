# Golang send USDC example

This example shows how to send USDC with MPCVault's API using Golang.

Before running this example, you need to execute the following:

```bash

# Install the protocol compiler plugins for Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Pull protobuf repository
git clone https://github.com/mpcvault/mpcvaultapis.git

# Generate gRPC code
protoc --go_out=./ --go-grpc_out=./ ./mpcvaultapis/mpcvault/platform/v1/*.proto
```