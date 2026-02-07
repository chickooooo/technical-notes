# Communication Protocols

<br>
<br>
<br>
<br>
<br>

## Index

- [TCP](#tcp)
- [UDP](#udp)
- [QUIC](#quic)
- [HTTP/1 vs HTTP/2 vs HTTP/3](#http1-vs-http2-vs-http3)
- [Polling (short & long)](#polling-short--long)

<br>
<br>
<br>
<br>
<br>

### TCP

- TCP stands for **Transmission Control Protocol**.
- It is a connection-oriented, reliable, transport layer protocol.
- It provides ordered delivery of data and also support re-transmission of lost packets.
- Common usecases: HTTP/HTTPS, FTP, databases, etc.
- **Note**: TCP provides a full duplex communication.

<br>

#### How TCP works?

Connection Establishment (3-Way Handshake)

- Client connects with the server using a **3-way handshake**.
    - `SYN`: Client requests to connect with the server.
    - `SYN + ACK`: Server acknowledges and requests to connect with the client.
    - `ACK`: Client acknowledges and connects with the server.
- A reliable & stateful connection is established between client-server.

---

Data Transfer

- TCP breaks data into **segments** and ensure:
    - **Sequencing**: Each segment has a sequence number and is sent in that order.
    - **Acknowledgement**: Sender sends acknowledgements when received data.
    - **Retransmission**: Lost segments are resent.
    - **Flow control**: Sender adapts to receiver's capacity. 
    - **Congestion control**: Sender adapts to network conditions.

---

Connection Termination

- After the data transfer is complete, the connection is terminated using a **4-way termination**.
    - `FIN`: Client requests to close the connection.
    - `ACK`: Server acknowledges this request.
    - `FIN`: Server requests to close the connection.
    - `ACK`: Client acknowledges this request.
- The connection is now closed from both the sides.

<br>

#### Advantages

- Provides a persistent, duplex connection.
- Guarantees data arrives completely and in the correct sequence.
- Detects lost or corrupted packets and retransmits them.
- Adapts to network conditions and avoids congestion.

<br>

#### Disadvantages

- Higher latency due to 3-way connection setup and retransmissions.
- Acknowledgements and state management adds extra overhead.
- A single lost packet can delay all following packets.
- Not ideal for real-time applications.

<br>
<br>
<br>
<br>

### UDP

- UDP stands for **User Datagram Protocol**.
- It is a connectionless, lightweight, transport layer protocol.
- It provides speed, low latency and minimal overhead during data transfer.
- Unlike TCP, it does not guarantee delivery, ordering, or duplication protection.
- Common usecases: Video/audio streaming (twitch), online gaming, live broadcasting, etc.

<br>

#### Key features

- No connection setup
    - There is no handshake (unlike TCP's 3-way handshake).
    - A sender can start sending data immediately.
- Datagram-based communication
    - Each message is sent as an independent packet (datagram).
    - Datagrams may arrive out of order, duplicated, or not at all.
- Minimal headers: UDP header is only 8 byte and contains:
    - Source port, Destination port
    - Length and Checksum
- Best-effort delivery
    - UDP relies on the IP layer for delivery.
    - No acknowledgments, retransmissions, or congestion control.
- Application responsibility
    - If reliability, ordering, or retries are needed, the application itself must implement them.
- Broadcast support
    - Can send one packet to multiple recipients.

<br>

#### Advantages

- It has lower latency than TCP, making it ideal for real-time applications.
- It is lightweight and fast due to small headers and less system resource requirements.
- Supports broadcasting and multicasting. Datagrams can be sent to multiple recipients at a time.
- Ideal for real-time applications. Lost packets does not need to be retransmitted.

<br>

#### Disadvantages

- No reliability: Packets can be lost, out-of-order, or corrupted.
- No Congestion Control: Can overwhelm the recipient network if not carefully managed.
- More Work for the Application: Application need to handle missing packets, ordering & retransmission.

<br>
<br>
<br>
<br>

### QUIC

- QUIC stands for Quick UDP Internet Connection.
- It is a connection-oriented, secure transport protocol that runs over UDP.
- Instead of relying on TCP for reliability and TLS for security, QUIC implements both itself in user space.
- HTTP/3 uses QUIC as its transport layer.

<br>

#### Key features & advantages

- Runs over UDP: Uses UDP as a base, but adds reliability, ordering and congestion control.
- Built-in Security: QUIC has TLS 1.3 encryption integrated into the protocol.
- Fast Connection Establishment: 1-RTT handshake for new connections. 0-RTT for repeat connections (client can send data immediately).
- Allows stream multiplexing without head-of-line blocking.
- Connection Migration:
    - Connections are identified by a Connection ID, not IP + port.
    - If IP changes (e.g., Wi-Fi to mobile data switch), the connection can continue. This is huge for mobile users.

<br>

#### Disadvantages

- Some firewalls, NATs, and proxies block or throttle UDP. Can cause fallback to TCP.
- More complex than TCP: reliability, security & congestion control all are handled in one protocol.
- User-space implementation and encryption increase CPU usage.
- 0-RTT data can be replayed by attackers. Requires careful application-level handling.

<br>
<br>
<br>
<br>

### HTTP/1 vs HTTP/2 vs HTTP/3

#### HTTP/1

- HTTP/1 uses plain text communication between the client and server, which is easy to read but inefficient.
- HTTP/1 creates a new TCP connection per request, which increases latency because of repeated handshakes. Also, it limits concurrency as a browser can only open around 6-8 connections for each server.
- HTTP/1.1 introduced persistent (keep-alive) connections which allows multiple requests to be sent through the same TCP connection. This introduced a new challenge of "Head of the line blocking". One slow request slows down all the other subsequent requests as these requests are processed sequentially.
- Not recommended. Upgrade to HTTP/2.

<br>

#### HTTP/2

- HTTP/2 uses binary representation for client - server communication. Which is more efficient for parsing and transmission.
- HTTP/2 supports **multiplexing**, allowing multiple requests and responses to be sent over a single TCP connection without waiting for each other. This eliminates the head-of-line blocking issue of HTTP/1.
- HTTP/2 uses HPACK (a compression algorithm) to reduce the size of HTTP headers, which saves bandwidth and reduce latency.
- HTTP/2 supports **server push**, which allows servers to send resources (like CSS, JS files) to the client proactively before the client even requests them, based on previous requests (HTML files). This can speed up page loads, but should be used cautiously.
- Recommended for all websites because most modern browsers support HTTP/2.

<br>

#### HTTP/3

- HTTP/3 is built on top of QUIC protocol, which uses UDP instead of TCP for communication. Similar to HTTP/2, communication happens using binary data.
- HTTP/3 uses more efficient handshake process which establishes connection faster and reduces latency. QUIC allows for **0-RTT** (zero round-trip time) connection establishment, meaning clients can begin sending data before the handshake is fully complete.
- QUIC natively encrypts traffic, making HTTP/3 more secure by default compared to previous versions that require additional SSL/TLS layers.
- Due to faster handshake, 0-RTT and low latency, HTTP/3 is more efficient for real-time applications like data streaming, video conference, online gaming, etc.
- Not widely used, but adoption is growing.

<br>
<br>
<br>
<br>

### Polling (short & long)

- In polling, a client repeatedly sends requests to the server at regular intervals to fetch data.
- Here, the client is responsible for initiating the request.
- The server responds with the current state, regardless of whether anything has changed since the last request.
- There are two types of polling: short and long.

<br>
<br>

#### Short polling

- In short polling, the client sends requests to the server at a fixed, frequent interval (e.g. every 5 seconds).
- The server responds immediately with whatever data it currently has, even if there are no new updates.

---

Key characteristics:

- Requests are sent on a fixed schedule.
- Responses are immediate.
- Many responses may contain no new data.
- Easy to implement and debug.

---

Drawbacks:

- Wastes bandwidth and server resources when updates are rare.
- Does not provide true real-time behavior.
- Can cause thundering-herd effects if many clients poll at the same time.

---

Use short polling when:

- Near-real-time updates are not required.
- The number of clients is relatively small.
- Simplicity is more important than efficiency.

<br>
<br>

#### Long polling

- In long polling, the client initiates the request and the server holds the request open until new data becomes available or a timeout occurs.
- Once the server responds, the client immediately sends a new request to wait for the next update.

---

Key characteristics:

- Requests stay open for a longer duration.
- The server responds only when data is available or a timeout is reached.
- Behavior approximates real-time updates.
- Fewer empty responses compared to short polling.

---

Drawbacks:

- More complex server and client side logic.
- Each connection costs server rousources. Not ideal for infrequent updates.

---

Use long polling when:

- Clients need near-real-time updates.
- WebSockets or SSE are not available or not feasible.

<br>
<br>

#### When to prefer polling over SSE & Websockets

- If the environment does not supports long-lived connections (SSE & websockets).
- If the frequency of updates is rare.
- If the number of clients is relatively low and a simple solution is required.

<br>
<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
