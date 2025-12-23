# Communication Protocols

<br>
<br>
<br>
<br>
<br>

## Index

- [TCP](#tcp)
    - [How TCP works?](#how-tcp-works)
    - [Advantages](#advantages)
    - [Disadvantages](#disadvantages)
- [UDP](#udp)
    - [Key features](#key-features)
    - [Advantages](#advantages-1)
    - [Disadvantages](#disadvantages-1)

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

### 

<br>
<br>
<br>
