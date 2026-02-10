# Database common concepts

<br>
<br>
<br>
<br>
<br>

## Index

- [Write ahead logging (WAL)](#write-ahead-logging-wal)

<br>
<br>
<br>
<br>
<br>

### Write ahead logging (WAL)

- Write ahead logging is a logging protocol where every change to the system is first written to a persistent log before the actual data is modified.
- The changes are first stored in a log and then applied to the main data store.
- WAL is needed to protect data against crashes, power failures, and partial writes.

<br>

#### Key features

- Durability: Once a transaction is committed, its changes are safely recorded in the log. Even if the system crashes immediately afterward, the data can be recovered from the log.
- Atomicity: WAL ensures that transactions are all-or-nothing. If a transaction fails halfway, the log allows the system to undo partial changes.
- Crash recovery: After a crash, the system can read the logs and can undo partial changes or can complete remaining changes.
- Sequential disk writes: Logs are written sequentially, which is much faster than random disk writes. This improves performance compared to writing directly to multiple data locations.

<br>

#### Drawbacks

- Double writes: Data is written at least twice; once to the log and once to the actual data store.
- Extra storage: The log consumes additional disk space and must be managed, archived, or truncated over time.
- Complexity: Introduces extra complexity when writing, recovering data, managing logs, etc.

<br>
<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
