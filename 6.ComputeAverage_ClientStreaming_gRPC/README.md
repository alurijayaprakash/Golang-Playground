# CompApp

### CompApp Implemented on top of gRPC-golang

## Client Streaming

## folder structure
----------------------------
### greetpb - contains proto buffer code - which is common for client and server
### greet-client - client code
### greet-server - server code


Parameters
1) -I: complete path of source directory
2) --go_out: set the output directory of the generated Go code
3) The last argument indicates the source file

How to install protobuf ?
https://developers.google.com/protocol-buffers/docs/downloads
choose win64X and update env variables too

Generate Protocode
1. pkg name in protofile should be like this
option go_package="/comppb";
2. generate protofile
-- to generate proto file only : protoc -I compApp/ --go_out=compApp/ compApp/comppb/comp.proto
-- to generate grpc and proto file : protoc -I compApp/ --go_out=compApp/ --go-grpc_out=compApp/ compApp/comppb/comp.proto


Go plugins for the protocol compiler:

https://grpc.io/docs/languages/go/quickstart/


Hint : go with grpc+proto

