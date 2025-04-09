import grpc
import json
import requests
from google.protobuf.json_format import MessageToDict
from google.protobuf.wrappers_pb2 import StringValue
import mpcvault.platform.v1.api_pb2 as pb
import mpcvault.platform.v1.api_pb2_grpc as pb_grpc

# ================= config =================
CONTRACT = "6656336A6018dD......."  
API_TOKEN = "your API TOKEN="
HTTP_ENDPOINT = "https://api.mpcvault.com/v1"
VAULT_UUID = "VAULT UUID"
CALLBACK_CLIENT_SIGNER_PUBLIC_KEY = "ssh-ed25519 ....."
# ============================================
# =================auth=======================
class GRPCAuth(grpc.AuthMetadataPlugin):
    def __init__(self, key):
        self._key = key

    def __call__(self, context, callback):
        callback((('x-mtoken', self._key),), None)

# ----------------- gRPC request -----------------
def grpc_call():
    try:
        channel = grpc.secure_channel(
            'api.mpcvault.com:443',
            grpc.composite_channel_credentials(
                grpc.ssl_channel_credentials(),
                grpc.metadata_call_credentials(GRPCAuth(API_TOKEN)),
            )
        )
        stub = pb_grpc.PlatformAPIStub(channel)
        evm_data_dict = {
            "chain_id": 137,
            "from": "0xc0abaa254729296a45a3885639AC7E10F9d54979",  # 直接使用 `from`
            "to": "",
            "value": "0",
            "input": bytes.fromhex(CONTRACT),
            "gas_fee": {
                "gas_limit": {"value": "10000000"}  # StringValue 的字典表示
            },
            "nonce": {"value": "0"}  # StringValue 的字典表示
        }

        evm_send_custom = pb.EVMSendCustom()

        request = pb.CreateSigningRequestRequest(
            evm_send_custom=evm_send_custom
        )

        response = stub.CreateSigningRequest(request)
        print("[gRPC] response:", response)
        return response
    except grpc.RpcError as e:
        print(f"[gRPC] failed: {e.code()}: {e.details()}")
    except Exception as e:
        print(f"[gRPC] error: {str(e)}")

# ----------------- HTTP request -----------------
def http_call():
    try:
        gas_fee = {
            "gas_limit": "10000000" 
        }

        evm_data = {
            "chain_id": 137,
            "from": "0xc0abaa254729296a45a3885639AC7E10F9d54979",
            "to": "",
            "value": "0",
            "input": bytes.fromhex(CONTRACT), 
            "gas_fee": gas_fee,
            "nonce": "0"  
        }
        request_body = {
            "evm_send_custom": evm_data  
        }
        if 'input' in request_body['evm_send_custom']:
            request_body['evm_send_custom']['input'] = "0x" + request_body['evm_send_custom']['input'].hex()
        headers = {
            "Content-Type": "application/json",
            "X-Mtoken": API_TOKEN
        }
        
        response = requests.post(
            f"{HTTP_ENDPOINT}/v1/createSigningRequest",
            headers=headers,
            data=json.dumps(request_body)
        )

        if response.status_code == 200:
            print("[HTTP] response:", response.json())
            return response.json()
        else:
            print(f"[HTTP] request failed {response.status_code}: {response.text}")
    except Exception as e:
        print(f"[HTTP] error: {str(e)}")

if __name__ == '__main__':
    print("========== do grpc request==========")
    grpc_result = grpc_call()
    
    print("\n========== do http request==========")
    http_result = http_call()
