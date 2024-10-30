# Nats (Beginner)

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
- Session - 1   : 1:30 hrs
- Tea Break     : 0:15 mins
- Session - 2   : 1:30 hrs
- Lunch Break   : 0:45 mins
- Session - 3   : 1:30 hrs
- Tea Break     : 0:15 mins
- Session - 4   : 1:15 hrs

## Methodology
- No powerpoint
- Discussion & Code
- No dedicated time for Q & A

## Prerequisites
- go tools

## About You
- Your Name
- Total Exp.
- Primary Language of software development
- Primary skillset

## What is NATS?
- Pub/Sub infrastructure


## Installation
- Reference (https://docs.nats.io/running-a-nats-service/introduction/installation)

- Using go tool
    - go install github.com/nats-io/nats-server/v2@latest
    - installed in the GOPATH/bin

## Communication with NATS
- telnet
- nats cli
- Language SDKs (go, python, c#, java)

## Using NATS CLI

---------------
Publishers
```
nats pub hello --count=-1 --sleep=1s "Publisher-1 Message #{{Count}} @{{TimeStamp}}"
nats pub hello --count=-1 --sleep=1s "Publisher-2 Message #{{Count}} @{{TimeStamp}}"
nats pub hello --count=-1 --sleep=1s "Publisher-3 Message #{{Count}} @{{TimeStamp}}"
```

Consumers (sharding)
```
nats sub hello --queue="hello-group"
nats sub hello --queue="hello-group"
```
-----------------