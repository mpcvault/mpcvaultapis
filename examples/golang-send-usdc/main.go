package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	apipb "go.mpcvault.com/go.mpcvault.com/genproto/mpcvaultapis/platform/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)

const APIEndpoint string = "api.mpcvault.com:443"
const APIToken = "[API Token]"
const VaultUUID = "[vault uuid]"
const CallbackClientSignerPublicKey = "[callback client signer public key]"

var grpcClient apipb.PlatformAPIClient

func main() {
	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.Dial(APIEndpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	header := metadata.New(map[string]string{
		"x-mtoken": APIToken,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	grpcClient = apipb.NewPlatformAPIClient(conn)

	// EVMSendERC20
	EVMSendERC20(ctx)
}

// EVMSendERC20 send erc20 token
func EVMSendERC20(ctx context.Context) {
	req := apipb.CreateSigningRequestRequest{}
	createSigningType := &apipb.CreateSigningRequestRequest_EvmSendErc20{}
	req.Type = createSigningType
	createSigningType.EvmSendErc20 = &apipb.EVMSendERC20{
		ChainId:              137,                                          // polygon chain id
		From:                 "0x71C7656EC7ab88b098defB751B7401B5f6d8976F", // sender address
		To:                   "0x71C7656EC7ab88b098defB751B7401B5f6d8976F", // receiver address
		TokenContractAddress: "0x2791bca1f2de4661ed88a30c99a7a9449aa84174", // USDC contract address on polygon
		Amount:               "1000000",                                    // 1 USDC
		GasFee:               nil,                                          // leave nil to use auto gas settings
		Nonce:                &wrappers.Int64Value{Value: 0},
	}
	req.Notes = &wrappers.StringValue{Value: "sending 1 USDC for testing"} // setting transaction notes

	res, err := grpcClient.CreateSigningRequest(ctx, &req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	b, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(b))
}

// EVMSendERC20 send erc20 token with client signer
func EVMSendERC20WithClientSigner(ctx context.Context) {
	req := apipb.CreateSigningRequestRequest{}
	createSigningType := &apipb.CreateSigningRequestRequest_EvmSendErc20{}
	req.Type = createSigningType
	createSigningType.EvmSendErc20 = &apipb.EVMSendERC20{
		ChainId:              137,                                          // polygon chain id
		From:                 "0x71C7656EC7ab88b098defB751B7401B5f6d8976F", // sender address
		To:                   "0x71C7656EC7ab88b098defB751B7401B5f6d8976F", // receiver address
		TokenContractAddress: "0x2791bca1f2de4661ed88a30c99a7a9449aa84174", // USDC contract address on polygon
		Amount:               "1000000",                                    // 1 USDC
		GasFee:               nil,                                          // leave nil to use auto gas settings
	}
	req.Notes = &wrappers.StringValue{Value: "sending 1 USDC for testing"} // setting transaction notes
	req.VaultUuid = &wrappers.StringValue{Value: VaultUUID}
	req.CallbackClientSignerPublicKey = &wrappers.StringValue{Value: CallbackClientSignerPublicKey}

	res, err := grpcClient.CreateSigningRequest(ctx, &req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	b, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(b))
}
