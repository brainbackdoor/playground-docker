### Use Docker-Compose

---
#### Docker Basic Commands

```
#### Container ID check
docker ps -a

#### Image Check
docker images

#### Log output
docker logs <container id>

#### Enter container
docker exec -it <container id> /bin/bash

#### docker container IP check
docker exec <container id> ip addr show eth0

#### docker image details
docker inspect <docker image name>

#### Docker All Container Stop
docker stop $(docker ps -a -q)

#### Docker All Container Remove
docker rm $(docker ps -a -q)

#### Docker All Images Remove
docker rmi $(docker images -q)
```

---


##### Image build
```
cd step04
docker build -t bbd/node-web-app .
docker build -t bbd/load-balance-nginx nginx/.
```

##### use Docker-compose
```
docker-compose up -d
```
