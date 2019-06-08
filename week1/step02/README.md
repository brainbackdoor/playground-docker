### Web Distribution

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
cd step02
docker build -t bbd/node-web-app .
```

##### Image execution
```
docker run -p 49160:8080 -d bbd/node-web-app
```
>-d로 이미지를 실행하면 분리 모드로 컨테이너를 실행해서 백그라운드에서 컨테이너가 돌아가도록 한다.  
-p 플래그는 공개 포트를 컨테이너 내의 비밀 포트로 리다이렉트한다.

##### Docker Machine IP check
```
docker-machine ip
```
##### Test
```
 curl -i <docker-machine ip>:49160
```
