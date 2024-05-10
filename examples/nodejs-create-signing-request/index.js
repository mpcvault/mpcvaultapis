// const express = require('express')
// const app = express()
// const port = 4202

// app.get('/', (req, res) => {
//     res.send('Hello World!');
//     console.log("Hello World");
// })

// let cors = require('cors');
// app.use(cors())

// app.listen(port, () => {
//     console.log(`Example app listening on port ${port}`)
// })

const API_TOKEN = "YOUR-API-TOKEN";
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const protobuf = require('protobufjs');
const protoPath = './proto/api.proto';
const packageDefinition = protoLoader.loadSync(protoPath, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const PlatformAPIService = protoDescriptor.mpcvault.platform.v1.PlatformAPI;
const client = new PlatformAPIService('api.mpcvault.com:443', grpc.credentials.createSsl());

// Create metadata 
const metadata = new grpc.Metadata();
metadata.add('x-mtoken', API_TOKEN);

const root = protobuf.loadSync(protoPath);
const EVMGas = root.lookupType('mpcvault.platform.v1.EVMGas');

// Create EVMGas
const gasObject = {
    maxFee: {value: '100000'}, // 使用google.protobuf.StringValue来设置值
    maxPriorityFee: {value: '50000'},
    gasLimit: {value: '21000'}
};
const evmGas = EVMGas.create(gasObject);

const CreateSigningRequestRequest = {
    evm_send_custom: {
        "chain_id": 137, // polygon
        "from": "0x544845005e42fE00a3C0E9735EEEC25Aa068b428", // your wallet address on MPCVault
        "to": "", // leave empty for contract creation
        "value": "0", // amount of tokens to send, in this case 0
        "input": [], // contract bytecode
        "gas_fee": evmGas, // you can leave this empty to use default gas settings
        "nonce": {value: '0'}, // you can leave this empty to use default nonce
    }
}
// send rpc request
client.CreateSigningRequest(CreateSigningRequestRequest, metadata, (error, response) => {
    if (error) {
        console.log('=============error==================');
        console.log(error);
    } else {
        console.log('=============response==================');
        console.log(response);
        // 使用 protobufjs 的 Message.decode 方法对响应进行反序列化
    }
});