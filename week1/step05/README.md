### Use Docker-Compose (3 Tier)

---
##### use Docker-compose
```
docker-compose -f ./docker-compose.yml up -d

docker ps
docker-compose stop
docker-compose rm
```

##### Wordpress 서버 IP 문제 해결

```
docker exec -it [데이터베이스컨테이너 이름] mysql -u wordpress -p
암호 입력 (기본값 : wordpress)
use wordpress
update wp_options set option_value='http://[서버 공인 IP 주소]' where option_name='siteurl' or option_name='home';
```