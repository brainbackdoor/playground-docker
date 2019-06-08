### Nginx(reverse-proxy) + Docker(node.JS)

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
cd step03
docker build -t bbd/node-web-app .
```

##### Image execution
```
docker run -e "MESSAGE=First instance" --name instance-1 -p 8081:8080 -d bbd/node-web-app 
docker run -e "MESSAGE=Second instance" --name instance-2 -p 8082:8080 -d bbd/node-web-app 
```

##### Docker Machine IP check
```
docker-machine ip
```
##### Test
```
curl -i <docker-machine ip>:8081
curl -i <docker-machine ip>:8082
```
---

### Build a Docker to Use as a Load Balancer(Nginx / reverse-proxy)

##### Image build
```
docker build -t bbd/load-balance-nginx nginx/.
```
##### Image execution
```
docker run -p 4000:80 --name lb -d bbd/load-balance-nginx
```
##### Test
```
curl -i 192.168.99.100:4000
curl -i 192.168.99.100:4000
...
```
