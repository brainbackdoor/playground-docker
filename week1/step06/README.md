### Use Docker-Compose (scalable nginx)

---
##### use Docker-compose

```
docker build -t node-server .
docker images
docker-compose -p practice up -d
docker ps

docker-compose -p practice stop
docker-compose -p practice rm
```

---
#### Docker Basic Commands
```
docker-compose -p practice scale web=10
docker logs practice_lb_1

```