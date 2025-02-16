How It Works
Genesis Block Creation:
The blockchain starts with a Genesis Block (first block).
Block Structure:
Each block contains an index, timestamp, data, previous hash, nonce, and its own hash.
Proof of Work (PoW):
The system finds a hash that starts with a certain number of zeros (difficulty level).
This simulates mining, ensuring that new blocks require computation.
Validation & Adding Blocks:
Before adding, blocks are checked for consistency and correctness.
Blockchain Printing:
After mining, the blockchain is printed with all its blocks.

Example Output


Index: 0
Timestamp: 2025-02-16 12:00:00 +0000 UTC
Data: Genesis Block
PrevHash:
Hash: 00005f9c8b1e3b1c4e2d879048f0e678ec4de1c1b3f38cfc879b8e5aee9e0e4b
Nonce: 0

Index: 1
Timestamp: 2025-02-16 12:01:23 +0000 UTC
Data: Alice sent 1 BTC to Bob
PrevHash: 00005f9c8b1e3b1c4e2d879048f0e678ec4de1c1b3f38cfc879b8e5aee9e0e4b
Hash: 0000a72fda1d4e3a1d64e67b7b5ec4a5e7d2c8f90d2b4f6e8b8c5d7f9c9e1b6e
Nonce: 5273

Index: 2
Timestamp: 2025-02-16 12:02:15 +0000 UTC
Data: Bob sent 0.5 BTC to Charlie
PrevHash: 0000a72fda1d4e3a1d64e67b7b5ec4a5e7d2c8f90d2b4f6e8b8c5d7f9c9e1b6e
Hash: 000073c2b8e4f89cde9a4f9a2c3f7d1e6f5a2b8c9e1d6e7b5c4d2a8f7c9b5e2d
Nonce: 13482

