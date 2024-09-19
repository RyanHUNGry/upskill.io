# Docker Startup

## Links
https://hub.docker.com/_/cassandra

Cassandra process is configurable through a runtime YAML file at path `/etc/cassandra/cassandra.yaml`. The underlying JVM is configured under a script at path `/etc/cassandra/cassandra-env.sh`. Base initialization of a Cassandra container uses ~8GB heap memory. 

The default JVM heap (object storage) allocation is:
- set max heap size based on the following
- max(min(1/2 ram, 1024MB), min(1/4 ram, 8GB))
- calculate 1/2 ram and cap to 1024MB
- calculate 1/4 ram and cap to 8192MB
- pick the max

This causes memory issues on a local machine running a multi-node setup, and ultimately leads to performance and container status issues. Cassandra is Java-based, and so the underlying JVM heap memory usage can be set as one of the variables in `/etc/cassandra/cassandra-env.sh`. The script uses settable environment variables (or manual editing).

Local machine allocates 8GB for docker, so try 1GB per container:
- MAX_HEAP_SIZE="1G"
- HEAP_NEWSIZE="400M"

Use a single-node setup for development. For a multi-node deployment, create a shared Docker network so CASSANDRA_SEEDS can use container hostname resolution.
