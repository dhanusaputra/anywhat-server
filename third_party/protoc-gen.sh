protoc --proto_path=api/pb --go_out=plugins=grpc:api/pb anywhat.proto
protoc --proto_path=api/pb --swagger_out=logtostderr=true:third_party/swagger-ui anywhat.proto