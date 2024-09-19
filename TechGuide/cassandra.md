Extracted from: https://cassandra.apache.org/doc/stable/cassandra/data_modeling/intro.html

# Partitions:
1. Cassandra is a distributed database that shards data using consistent hashing and a partition key
2. A primary key can be any pair of (x, y, v, ..., j), and acts similar to SQL
    a. It can also be ((x, y, ..., v), j, x, ..., z), where the parantheses indicate partition key
3. The first component of the primary key is the partition key, which is used for consistent hashing and sharding
4. The remaining components of the primary key are cluster keys, which determine order in which data is stored within a partition

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