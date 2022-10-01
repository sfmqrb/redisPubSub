# redisPubSub
dockerized goroutine subscribers consume data published by python process using redis 


Contains two projects and processes in production one written in golang and the other in python.

## Python process is the publisher.
It simulates one threaded application to generate data into queue.

## Golang process is subscriber.
The other process's code is written in golang.
Using green threads of go, It consumes data which was published.

In this application I use redis Pub/Sub to notify Subs that there is actually something new in the queue and then LPOP that.
So instead of spinning around the LIST we use this mechanism to avoid polling.

[more info about push-based redis pub/sub](https://blog.devgenius.io/how-to-use-redis-pub-sub-in-your-python-application-b6d5e11fc8de)

## run

```bash
python3 publisher.py
```


```bash
go run subscriber.go
```

make sure redis-server is running on default port.

## todo
use docker-compose
