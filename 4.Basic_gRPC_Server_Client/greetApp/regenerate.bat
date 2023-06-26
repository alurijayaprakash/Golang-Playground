go get -u google.golang.org/grpc

go get -u github.com/golang/protobuf/protoc-gen-go

protoc greetApp/greetpb/greet.proto --go_out=plugins=grpc:.
----------------------------
working : protoc -I=greetApp/ --go_out=greetApp/ greetApp/greet.proto

option go_package="/greetpb";
protoc -I=greetApp/greetpb/ --go_out=greetApp/greetpb/ greetApp/greetpb/greet.proto





protoc -I greetApp/ --go_out=greetApp/ greetApp/greetpb/greet.proto