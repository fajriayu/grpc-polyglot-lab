
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = '../user.proto';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: false,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const userProto = grpc.loadPackageDefinition(packageDefinition).user;

const client = new userProto.UserService(
    'localhost:50051',
    grpc.credentials.createInsecure()
);

client.GetUserProfile({ userId: 1 }, (error, response) => {
    if (error) {
        console.error('Gagal mengambil data:', error.message);
        return;
    }

console.log('Nama  :', response.name);
console.log('Email :', response.email);
console.log('Phone :', response.phone);
});
