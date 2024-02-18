# DCServer
 Yandex Lyceum Project - Distributed Computing Server

Step 1:


# Using examples:
 Send expression:
 ```bash
 curl -X POST "expression=2p2*2" http://localhost:8080/tasks
```
 Check answer:
 ```bash
curl http://localhost:8080/tasks?id=12
```
