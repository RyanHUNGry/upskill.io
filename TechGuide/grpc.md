Extracted from: https://grpc.io/docs/

# What is gRPC?:
1. gRPC is an RPC framework
2. Allows clients to call methods defined by an interface, which is implemented by a server
3. The service contract (or interface) is established through ProtoBuf
    a. ProtoBuf helps serialize structured data
    b. Helps generate ProtoBuf serialization and deserialization code through protoc, and protoco has a gRPC plugin to help generate client and server stubs used in conjunction with serialization and deserlization code
4. Suitable for microservice communication, and decouples implementation and calling details to client stubs and server skeleton or stubs

# RDBMS vs Cassandra:
1. Cassandra does not support joins, and recommends denormalized data
2. Cassandra does not enforce referential integrity
3. Building a Cassandra data model starts from structuring tables around queries, which is opposite in nature to RDBMS
4. Ordering is a design decision in Cassandra, whereas it is a query feature in RDBMS

# Logical Data Modeling:
1. After defining a set of queries, create a graph where query results can unlock downstream queries
    a. This is handled at the application level, where multiple queries build on top of eachother unlike at the DB level with SQL JOINs
    b. Write this kind of logic in application code because Cassandra does not support multi-table queries and JOINs
2. For each query, figure out what the query is searching by, and utilize this criteria as the partition key
3. For each query, keep the details of its corresponding table to be as minimal as possible to answer the query, but this is just a design decision

# Resources:
1. https://docs.datastax.com/en/cassandra-oss/3.x/cassandra/cassandraAbout.html