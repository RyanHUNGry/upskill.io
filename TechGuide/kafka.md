Extracted from: https://developer.confluent.io/courses/apache-kafka/events/

# Topics:
1. Unit of event organization
2. Similar to a table in a relational database, where similar entities or events are stored together
3. A topic is not a queue (though some might use it as an umbrella term) but rather a set of logs
    a. Logs are append-only data structures that can only be seeked by offset rather than through indexing
    b. Events are immutable within a topic, making it easier to replicate across nodes
    c. Topics are durable because events stored can only be deleted through a configurable retention period
    d. Underlying logs are stored on disk rather than in-memory, making data durable
4. If a topic has one only partition, then the topic is the sole log

# Partitions:
1. Kafka is a distrbuted system
2. A topic on a single-node system can only scale to a limiting extent, so topics can be partitioned into multiple logs
3. Events are sent to a certain partition using consistent hashing if a key exists, or round-robin if not
4. Partitions locally preserve the order of messages that share a key

Brokers:
1. A computer, instance, container, or VM running Kafka process
2. Manage partitions, partition replication, and handle read and write requests

# Replication:
1. Copies of data for fault tolerance
2. Replication is done on the partition-level under a leader-follower architecture N - 1 followers moderate a lead partition
    a. Leader-follower differs from master-slave in some aspects
3. Read and writes go are handled by the lead partition
4. When the leader fails, a follow is elected as the new cluster leader

# Producers:
1. At a high level, the producer API exposes a producer and event object
2. Partitioning is performed by the producer, which then sends the message to the correct partition
