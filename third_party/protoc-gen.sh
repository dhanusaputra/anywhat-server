protoc --proto_path=api/pb --go_out=plugins=grpc:api/pb anywhat.proto
protoc --proto_path=api/pb --go_out=plugins=grpc:api/pb user.proto
protoc --proto_path=api/pb --swagger_out=logtostderr=true:third_party/swagger-ui anywhat.proto
protoc --proto_path=api/pb --swagger_out=logtostderr=true:third_party/swagger-ui user.proto
protoc --proto_path=api/pb --swagger_out=logtostderr=true:api/swagger anywhat.proto
protoc --proto_path=api/pb --swagger_out=logtostderr=true:api/swagger user.proto