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

### Namespacing
Publishers
```
nats pub org.order.submitted --count=-1 --sleep=2s "Order {{Count}} Submitted"
nats pub org.order.processed --count=-1 --sleep=2s "Order {{Count}} Processed"
nats pub org.payment.submitted --count=-1 --sleep=2s "Payment {{Count}} Submitted"
nats pub org.payment.processed --count=-1 --sleep=2s "Payment {{Count}} Processed"
```

Subscribers
```
nats sub ">"
nats sub "org.>"
nats sub "org.order.>"
nats sub "org.payment.>"
nats sub "org.*.submitted"
nats sub "org.*.processed"
nats sub "org.*.*.submitted"
```

## Request / Reply
```
nats reply foo "service instance A Reply# {{Count}}"
```
```
nats request foo --count 2 "Request {{Count}}"
```

### Cluster
```
nats-server -n "server-1" -D -p 4222 -cluster nats://localhost:6222 --cluster_name=my_cluster --routes nats://localhost:6333,nats://localhost:6444

nats-server -n "server-2" -D -p 4333 -cluster nats://localhost:6333 --cluster_name=my_cluster --routes nats://localhost:6222,nats://localhost:6444

nats-server -n "server-3" -D -p 4444 -cluster nats://localhost:6444 --cluster_name=my_cluster --routes nats://localhost:6222,nats://localhost:6333
```

#### Cluster Auto Discovery

```
nats-server -n "server-1" -D -p 4222 -cluster nats://localhost:6222 --cluster_name=my_cluster 

nats-server -n "server-2" -D -p 4333 -cluster nats://localhost:6333 --cluster_name=my_cluster --routes nats://localhost:6222

nats-server -n "server-3" -D -p 4444 -cluster nats://localhost:6444 --cluster_name=my_cluster --routes nats://localhost:6222
```

## Shutdown server
```
nats-server -sl quit=<pid>




