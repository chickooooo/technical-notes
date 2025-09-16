**How Hashing works?**

Hash Function:
- A normal function that must follow these 3 rules:
    - **Deterministic**: Same input should always produce the same output.
    - **Fixed length output**: For all possible inputs, the length of the output should be same.
    - **Irreversible**: It should be impossible to reverse engineer the input from the output.

Hash:
- Output of a hash function is called a **Hash**.

Hashing:
- The process of converting a **Key** into a hash is called **Hashing** 

Hash Collision:
- When two or more inputs (Keys) to a hash function returns the same output (Hash) it is called **Hash Collision**.
- A perfect hash function has no collisions at all.
- A good hash function tries to minimise hash collisions.

<br>
<br>

**Encoding**

Encoding:
- Encoding transforms data into a different format using a known scheme, so that it can be safely stored or transmitted and then correctly **decoded back**.
- This process is fully reversible.
- Encoding is not meant to product data from unauthorised access.
- It is used when storing or transmitting data.
- Use cases:
    - Encoding URLs (convert `space` to `%20`).
    - Storing non-ASCII characters in JSON / database.

- It is reversible. Two way.
- About data representation and not security.
- e.g. CSS color hex (#303030) to RGB (48, 48, 48).

<br>
<br>

**Encryption**

Encryption:
- Encryption is the process of converting plain text into **cipher text** using a **key**, so that only authorised parties (with the correct key) can decrypt and access the original data.
- This process is reversible with the correct key.
- It is used to protect the data from unauthorised access.
- Types of encryption:
    - Symmetric Encryption: Same key is used for encryption and decryption.
    - Asymmetric Encryption: Uses a **public** key for encryption and a **private** key for decryption.
- Use cases:
    - HTTPS (secure web communication).
    - Secure file storage.
    - Whatsapp message encryption.

- A key is needed to reverse back to original data.
- e.g. whats app message encryption.

<br>
<br>

**Hashing**

Hashing:
- Hashing is the process of converting data of any size into a **fixed-size string**.
- The function used is called a **Hash function** and the output is called a **Hash**.
- This process is non-reversible. It is a one-way conversion.
- It is used to verify that the data has not been altered.
- Use cases:
    - Password storage (store hash, not actual password)
    - Data integrity checks (checksums)
    - Hash tables (in-memory key-value lookups)
    - Distributed systems (consistent hashing for sharding)

<br>
<br>
<br>
<br>
<br>

- Popular hashing algorithms
- Popular encryption algorithms
- Popular encoding algorithms
- checksum

- consistent hashing
