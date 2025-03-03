# protoc usage
protoc can operate via full path via CWD or relative path given your location in shell

protoc --proto_path=<path to directory containing .proto file> --go_out=<output directory> --go_opt=paths=source_relative <name of the proto file>

example in grpc directory:
protoc --proto_path=. --go_out=. .proto