package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	apipb "go.mpcvault.com/genproto/mpcvault/platform/v1"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

const APIEndpoint string = "api.mpcvault.com:443"
const HttpEndpoint string = "https://api.mpcvault.com/v1"
const APIToken = "your API TOKEN"
const VaultUUID = "your vault uuid"
const CallbackClientSignerPublicKey = "[callback client signer public key]"

var (
	wg         sync.WaitGroup
	grpcClient apipb.PlatformAPIClient
)

func main() {
	wg.Add(2)

	go func() {
		defer wg.Done()
		httpDemo()

	}()
	go func() {
		defer wg.Done()
		grpcDemo()
	}()
	wg.Wait()
}

func httpDemo() {
	method := "createSigningRequest"
	url := fmt.Sprintf("%s/%s", HttpEndpoint, method)
	body, err := protojson.MarshalOptions{
		UseProtoNames: true,
	}.Marshal(creatRequestBody())
	if err != nil {
		fmt.Println("HTTP RESULT:", err)
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	//add header
	req.Header.Set("X-Mtoken", APIToken)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP RESULT:", err)
	} else {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("HTTP Status:", resp.Status)
		fmt.Println("Response Body:", string(bodyBytes))
	}
	defer resp.Body.Close()
}

func grpcDemo() {
	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.Dial(APIEndpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("GRPC RESULT:", err)
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
	req := creatRequestBody()
	res, err := grpcClient.CreateSigningRequest(ctx, req)
	if err != nil {
		fmt.Println("GRPC RESULT:", err)
	} else {
		b, _ := json.MarshalIndent(res, "", " ")
		fmt.Println("GRPC RESULT:", string(b))
	}
}

func creatRequestBody() *apipb.CreateSigningRequestRequest {
	req := &apipb.CreateSigningRequestRequest{}
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
	return req
}

// EVMSendERC20 send erc20 token with client signer
func EVMSendERC20WithClientSigner(ctx context.Context) {
	req := creatRequestBody()
	req.VaultUuid = &wrappers.StringValue{Value: VaultUUID}
	req.CallbackClientSignerPublicKey = &wrappers.StringValue{Value: CallbackClientSignerPublicKey}
	res, err := grpcClient.CreateSigningRequest(ctx, req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	b, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(b))
}
