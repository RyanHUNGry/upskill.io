Extracted from: https://grpc.io/docs/

# What is gRPC?:
1. gRPC is an RPC framework
2. Allows clients to call methods defined by an interface, which is implemented by a server
3. The service contract (or interface) is established through ProtoBuf
    a. ProtoBuf helps serialize structured data
    b. Helps generate ProtoBuf serialization and deserialization code through protoc, and protoco has a gRPC plugin to help generate client and server stubs used in conjunction with serialization and deserlization code
4. Suitable for microservice communication, and decouples implementation and calling details to client stubs and server skeleton or stubs
