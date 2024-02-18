# DCServer
 # Yandex Lyceum Project - Distributed Computing Server
 ![Logo](https://github.com/Mendium/DCServerYa/blob/main/orig.png)

## Setup:
 (1) Install [MinGW64 by instruction](https://www.msys2.org/), and set path (C:\msys64\mingw64\bin  -- by default) to %PATH% (global variable):
 P/S: If folder /bin is empty download mingw64 and insert manually to this folder (or dm me).

 (2) Install all dependencies:
```bash
go mod download
```
 (3) Update CGO:
```bash
go env -w CGO_ENABLED=1
```
 (4): download this project (.zip), open in IDE (i recommend GoLand) and start the server (main.go).

 (5): Send your first expression to server by cmd :)

## If you have any questions/problems:
## My contacts: tg: @rielexx, discord: darklexx_

## Sign "+"  --> "p" (like in example)

# Using examples:
 ## Send expression:
 ```bash
 curl -X POST "expression=2p2*2" http://localhost:8080/tasks
```
Success:
```bash
Task added to database with ID: 3
```

 ## Check answer:
 ```bash
curl http://localhost:8080/tasks?id=12
```
Answer:
```bash
Your expression has been calculated. Result: 0
```
 ## Information about operations:
```bash
curl http://localhost:8080/operations
```

## How to write negative numbers:
```bash
-3  ---> (0-1)
-12 ---> (0-12)
```

# How does it works?
![Scheme](https://github.com/Mendium/DCServerYa/blob/main/info.png)
