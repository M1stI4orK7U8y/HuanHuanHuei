protobuf build :
cd /path/to/huanhuanhuei
protoc -I. --go_out=plugins=grpc:. src\database\types\record\record.proto