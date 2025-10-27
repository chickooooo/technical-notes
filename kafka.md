# Apache Kafka

<br>
<br>
<br>
<br>
<br>

## Index

<br>
<br>
<br>
<br>
<br>

### What is Kafka?

- Apache Kafka is a distributed event streaming platform.
- It is used for building real-time data streaming applications.
- It is a **high throughput**, **fault tolerant**, publish-subscribe messaging system.
- It allows services to communicate via sending messages (events) through **topics** in real-time.

<br>
<br>
<br>

### Key Components:

#### Producer

- Sends (publishes) messages to Kafka topics.

<br>

#### Consumer

- Reads (subscribes to) messages from a Kafka topics.
- A consumer can belong to a consumer group for load balancing.
- Each partition can be consumed by only one consumer from a consumer group.

<br>

#### Consumer group

- A group of consumers that consumes from partitions of a topic.
- They ensure load balancing and parallel processing.
- A consumer group gaurantees that each message is processed exactly once by the group.

<br>

#### Topic

- Logical channel or category where messages are published.
- A topic can have multiple partitions for scalability and parallelism.
- Events in a topic are **ordered within a partition** but not necessarily across partitions.

<br>

#### Partition

- A topic is split into partitions, which are units of parallelism.
- Each partition is an **ordered immutable** sequence of messages.
- Messages are identified by a **unique offset**.
- This offset enables fault tolerance.

<br>

#### Replication

- Each partition has multiple replicas across brokers.
- One broker hosts the leader replica (handles read/write) and others host follower replica (for fault tolerance).

<br>

#### Broker

- A broker is just a Kafka server that stores and serves messages.
- It receives messages from the producers and makes them available to the consumers.
- Messages are presisted on the disk and replicated for fault tolerance.

<br>

#### Cluster

- A cluster is a collection of one or more kafka brokers.
- Multiple brokers are used for scalability and fault tolerance.

<br>

#### ZooKeeper / Kafka Raft (KRaft)

- Manages cluster metadata and leader election.
- Keeps track of brokers, topics, partitions and their leaders.
- Ensures fault tolerance.

<br>
<br>
<br>

### 

<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
<br>

- Databases have low throughput
- Kafka has high throughput
- Kafka is a hybrid combination of queue + pub-sub model.
- Multiple consumer groups allow kafka to behave as a pub-sub model.

<br>

- Producer: produce data
- Consumer: consume data
- Consumers consume from kafka and bulk-insert into the database

<br>

- Topic:
    - Logical partitioning of the messages.
    - Example: order, cart, etc.
    - A topic is further divided into partitions.

<br>

- Partition:
    - A producer produces a message into a partition.
    - A consumer consumes messages from a partition.

<br>

```
+----------------------------------------------------+
|                     KAFKA                          |
|                                                    |
|   +--------------------+    +--------------------+ |
|   |     TOPIC 1        |    |     TOPIC 2        | |
|   |                    |    |                    | |
|   |  +-------------+   |    |                    | |
|   |  | PARTITION 1 |   |    |                    | |
|   |  +-------------+   |    |                    | |
|   |  +-------------+   |    |                    | |
|   |  | PARTITION 2 |   |    |                    | |
|   |  +-------------+   |    |                    | |
|   +--------------------+    +--------------------+ |
|                                                    |
+----------------------------------------------------+

``` 

<br>

- Suppose we have 1 topic, this topic has 4 partitions and we have **1** consumer.
    - Then this consumer will consume from all 4 partitions.
    - Example: `C1 -> P1,P2,P3,P4`.
- Suppose we have 1 topic, this topic has 4 partitions and we have **2** consumer.
    - Then both consumers will consume from 2 partitions each.
    - Example: `C1 -> P1,P2` and `C2 -> P3,P4`.
- Suppose we have 1 topic, this topic has 4 partitions and we have **3** consumer.
    - Then 1 consumer will consume from 2 partitions and other 2 will consume from 1 partition each.
    - Example: `C1 -> P1,P2` and `C2 -> P3` and `C3 -> P4`.
- Suppose we have 1 topic, this topic has 4 partitions and we have **4** consumer.
    - Then each consumer will consume from 1 partition each.
    - Example: `C1 -> P1` and `C2 -> P2` and `C3 -> P3` and `C4 -> P4`.
- Suppose we have 1 topic, this topic has 4 partitions and we have **5** consumer.
    - Then 4 consumers will consume from 1 partition each and the remaining consumer will not consume from any partition.
    - Example: `C1 -> P1` and `C2 -> P2` and `C3 -> P3` and `C4 -> P4` and `C5 -> -`.

<br>

- A consumer can consume from multiple partitions.
- A partition can be consumed by only 1 consumer from a consumer group.
- A partition can be consumed by 2 consumers from 2 different consumer groups.

<br>

- Consumer group:
    - A group of consumers, all doing the same task.
    - Each consumer belongs to exactly 1 consumer group.
    - Example: Analytics consumer group, Invoice consumer group, etc.

<br>

- 
