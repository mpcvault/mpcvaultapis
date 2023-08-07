import grpc
import mpcvaultapis.mpcvault.platform.v1.api_pb2 as pb
import mpcvaultapis.mpcvault.platform.v1.api_pb2_grpc as pb_grpc
from google.protobuf.wrappers_pb2 import StringValue

CONTRACT = "79204861726468617420546f6b656e0............" # hex encoded contract bytecode
API_TOKEN = "[YOUR API TOKEN]"

class GRPCAuth(grpc.AuthMetadataPlugin):
    def __init__(self, key):
        self._key = key

    def __call__(self, context, callback):
        callback((('x-mtoken', self._key),), None)

if __name__ == '__main__':
    channel = grpc.secure_channel('api.mpcvault.com:443', 
                                  grpc.composite_channel_credentials(
                                      grpc.ssl_channel_credentials(),
                                      grpc.metadata_call_credentials(GRPCAuth(API_TOKEN)),
                                  ))
    stub = pb_grpc.PlatformAPIStub(channel)

    gas_fee = pb.EVMGas(gas_limit=StringValue(value="10000000")) # you can leave other fields empty to use default gas settings

    response = stub.CreateSigningRequest(pb.CreateSigningRequestRequest(evm_send_custom = pb.EVMSendCustom(**{
            "chain_id":137, #polygon
            "from":"0xc0abaa254729296a45a3885639AC7E10F9d54979", # your wallet address on MPCVault
            "to":"", # leave empty for contract creation 
            "value":"0", # amount of tokens to send, in this case 0
            "input":bytes.fromhex(CONTRACT), # contract bytecode
            "gas_fee":gas_fee, # you can leave this empty to use default gas settings
        })))
    
    print("Received message: ", response)