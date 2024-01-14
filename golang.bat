::golang
protoc -I=./ ^
 --go_out=C:\_workspace ^
 protofile/proto/types/*.proto ^
 protofile/proto/filesystem/*.proto ^
 protofile/proto/version/*.proto 
 


protoc -I=./ ^
 --go-grpc_out=C:\_workspace ^
 protofile/proto/filesystem/1grpc.proto ^
 protofile/proto/version/*.proto 