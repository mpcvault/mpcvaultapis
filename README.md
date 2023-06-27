# MPCVault APIs

This repo contains public interface definitions of MPCVault APIs.

MPCVault APIs use Protocol Buffers version 3 (proto3) as the Interface Definition Language (IDL) to define the API interface and the structure of the payload messages.

You can access MPCVault APIs published in this repository through GRPC, which is a high-performance binary RPC protocol over HTTP/2. It offers many useful features, including request/response multiplex and full-duplex streaming.

## Generate gRPC Source Code and Client Libraries

To generate gRPC source code for Google APIs in this repository, you first need to install both Protocol Buffers and gRPC on your local machine

We recommend looking through the official tutorials [here](https://grpc.io/docs/languages) to get started with gRPC.
