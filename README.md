![img.png](img.png)

## [Flowchart]

```mermaid
flowchart LR
    subgraph Your Backend
        B1[User_1 Profile]
        B2[User_1 Profile]
        B3[User_2 Profile]
    end

    subgraph Hicon Query Proxy
        direction TB
        combine[Combine Identical Queries]
        execute[Execute Single Identical Queries]
    end

    DB[(Redis\nor SQL Database...)]
    B1 & B2 --> combine
    B3 --> combine
    combine --> execute
    execute <-->|User_1 Profile| DB
    execute <-->|User_2 Profile| DB
    execute -->|Shared Result| B1
    execute -->|Shared Result| B2
    execute -->|Separate Result| B3
```

## Core Features

- [x] Built-in query builder use BunORM for security.
- [x] Combine identical queries into single query in the same time.
- [x] Cache connections in pool.
- [x] Support for multi DBs: MySQL, Postgresql.
- [x] Connect with your Redis for better performance.
- [x] Optimized Redis cache size memory if using built-in query builder.
- [x] Client SDK for NodeJS, Golang.
- [x] Debug logging with your X-Request-Id.
- [x] Disable cache at global and per request.
- [x] Custom your lock key with write actions.
- [x] Bulk write operations with transactions.

## Backlogs

- [ ] Support for Oracle.
- [ ] Unit test.
- [ ] SDK for PHP, C#, Ruby...
- [ ] OpenTelemetry.
