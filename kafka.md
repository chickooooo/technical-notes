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
