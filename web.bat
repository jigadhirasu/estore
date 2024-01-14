::productpb

del "jspb" /S /Q

protoc -I=./ ^
    --js_out=import_style=commonjs:./jspb ^
    --grpc-web_out=import_style=typescript,mode=grpcweb:./jspb ^
    protofile/proto/version/*.proto ^
    protofile/proto/resource/*.proto ^
    protofile/proto/types/*.proto ^
    protofile/grpc/*.proto
