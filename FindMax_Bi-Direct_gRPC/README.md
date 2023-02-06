# FindMaxApp

### FindMaxApp Implemented on top of gRPC-golang
## Bi-Directional Streaming


## folder structure
----------------------------
### findmaxpb - contains proto buffer code - which is common for client and server
### findmax-client - client code
### findmax-server - server code


Parameters
1) -I: complete path of source directory
2) --go_out: set the output directory of the generated Go code
3) The last argument indicates the source file

How to install protobuf ?
https://developers.google.com/protocol-buffers/docs/downloads
choose win64X and update env variables too

Generate Protocode
1. pkg name in protofile should be like this
option go_package="/greetpb";
2. generate protofile
-- to generate proto file only : protoc -I FindMaxApp/ --go_out=FindMaxApp/ FindMaxApp/findmaxpb/greet.proto
-- to generate grpc and proto file : protoc -I FindMaxApp/ --go_out=FindMaxApp/ --go-grpc_out=FindMaxApp/ FindMaxApp/findmaxpb/findmax.proto



Go plugins for the protocol compiler:

https://grpc.io/docs/languages/go/quickstart/


Hint : go with grpc+proto

