## 1. 도커란 무엇인가

>도커는 컨테이너형 가상화 기술을 구현하기 위한 상주 애플리케이션과 이 애플리케이션을 조작하기 위한 명령행 도구로 구성되는 프로덕트이다. 

#### 주요 특징

```
1. 재현성
 - 버전관리( + 이미지 내용을 코드화) 
   다양한 의존 라이브러리나 도구를 도커 컨테이너에 포함시켜 배포함으로써 실행 환경과 상관없이 스크립트의 동작 재현성을 높임

2. 가볍다

3. 이식성 (AUFS, overlay2) 
 - 개발환경과 운영환경을 거의 동등하게 재현
 
4. 비용절감

5. 표준화
```

- 도커의 기본 개념
```
- 컨테이너형 가상화 기술(LXC-> runC)
    LXC는 호스트 운영체제 가상화보다 성능면에서 유리하다는 장점으로 시스템 컨테이너로서 어느정도 자리를 잡았다. 

- 도커와 LXC 차이

    1. 호스트 운영 체제의 영향을 ㅂ다지 않는 실행환경(docker engine을 이용한 실행환경 표준화)
    2. DSL(Dockerfile)을 이용한 컨테이너 구성 및 애플리케이션 배포 정의
    3. 이미지 버전 관리
    4. 레이어 구조를 갖는 이미지 포맷(차분 빌드가 가능함)
    5. 도커 레지스트리(이미지 저장 서버 역할을 함)
    6. 프로그램 가능한 다양한 기능의 API
    
--> 도커는 컨테이너에 애플리케이션 실행환경이 함께 배포되는 방식이다. || 의존성 문제를 해결    
```

- 도커 스타일 체험하기
```
FROM절 : 컨테이너의 원형(틀) 역할을 할 도커 이미지(운영체제)를 정의
COPY절: 특정 파일을 도커 컨테이너 안의 특정 경로에 복사
RUN 절: 도커 컨테이너 안에서 어떤 명령을 수행하기 위한 것
CMD절: 완성된 이미지를 도커 컨테이너로 실행하기 전에 먼저 실행할 명령을 정의
```
## 2. 도커를 사용하는 의의

#### 1. 변화하지 않는 실행환경으로 멱등성 확보
#### 2. 코드를 통한 실행환경 구축 및 애플리케이션 구성
#### 3. 실행 환경과 애플리케이션의 일체화로 이식성 향상
#### 4. 시스템을 구성하는 애플리케이션 및 미들웨어의 관리 용이성


- 환경 차이로 인한 문제 방지
```
각 서버에 배포된 애플리케이션이 동일하다면 애플리케이션이 의존하는 환경의 차이를 가능한 한 배제하여야 한다. 
--> 코드로 관리하는 인프라와 불변 인프라

코드 기반으로 인프라 구축을 관리한다고 해도 멱등성을 보장하기 위해 항구적인 코드를 계속 작성하는 것은 운영 업무에 부담을 주기 쉽다. 
서버의 대수가 늘어날수록 모든 서버에 구성을 적용하는 시간도 늘어난다. 

불변인프라는 어떤 시점의 서버 상태를 저장해 복제할 수 있게하자는 개념이다.
서버에 변경을 가하고 싶은 경우에는 기존 인프라를 수정하는 대신 새로운 서버를 구축하고 그 상태를 이미지로 저장한 다음 그 이미지를 복제한다.

- 애플리케이션 구성 관리의 용이성
- 운영환경에서 빛을 발하는 도커
- 새로운 개발 스타일
```

---

## 2. 도커 컨테이너 배포

1. 컨테이너로 애플리케이션 실행하기
```
- 도커 이미지와 도커 컨테이너
1. 도커이미지: 도커 컨테이너를 구성하는 파일시스템과 실행할 애플리케이션 설정을 하나로 합친 것으로, 컨테이너를 생성하는 템플릿 역할을 한다.
2. 도커 컨테이너: 도커 이미지를 기반으로 생성되며, 파일시스템과 애플리케이션이 구체화되어 실행되는 상태
```


#### step2: http web server, port forwarding
```
docker image pull gihyodocker/echo:latest
docker container run -t -p 9000:8080 gihyodocker/echo:latest
curl http://localhost:9000
docker stop $(docker container ls -q)
```

- 간단한 애플리케이션과 도커 이미지 만들기
```
FROM 인스트럭션: 도커 이미지의 바탕이 될 베이스 이미지를 지정한다.
RUN 인스트럭션: 도커 이미지를 실행할 때 컨테이너 안에서 실행할 명령을 정의
COPY 인스트럭션: 도커가 동작중인 호스트머신의 파일이나 디렉터리를 도커 컨테이너 안으로 복사하는 인스트럭션
CMD 인스트럭션: 도커 컨테이너를 실행할 때 컨테이너 안에서 실행할 프로세스를 지정한다. RUN은 이미지를 빌드할 떄 실행되고 CMD는 컨테이너를 시작할 떄 한번 실행된다.
```

#### step3
```
- 도커 이미지 빌드하기
- 도커 컨테이너 실행

컨테이너 포트 --> 포트포워딩이 필요
docker image build -t example/echo:latest .
docker container run -d -p 9000:8080 example/echo:latest
docker stop $(docker container ls --filter "ancestor=example/echo" -q
```
#### 2. 도커 이미지 다루기

```
도커 컨테이너를 만들기 위한 템플릿
- docker image build
- docker search 
- docker image pull
- docker image ls
- docker image tag
- docker image push
docker login -u brainbackdoor -p
docker image tag example/echo brainbackdoor/echo:latest
docker image push brainbackdoor/echo:latest
```

#### 3. 도커 컨테이너 다루기
```
- 도커 컨테이너의 생애주기
// 실행중상태, 정지상태, 파기상태
```

```
- docker container run 
docker container run -it alpine:3.7 uname -a
docker container run -t -d --name bbd8855 example/echo:latest

    -i : 컨테이너쪽 표준입력과의 연결을 그대로 유지
    -t : 유사터미널 기능을 활성화
    --rm : 컨테이너를 종료할 때 컨테이너를 파기하도록 하는 옵션
    -v : 호스트와 컨테이너간 디렉터리 혹은 파일을 공유하기 위한 옵션
```

- docker container ls
```
원하는 이름과 컨테이너명이 일치하는지 여부 --filter "name=echo1"
생성한 이미지를 기준으로 --filter "ancestor=example/echo"
```

- docker container stop
- docker container restart
- docker container rm (--rm 옵션)

- docker container logs
```
docker container logs -f $(docker container ls --filter "name=bbd8855" -q)
```

- docker container exec
```
컨테이너 내부의 상태를 확인하거나 디버깅하는 용도로 사용
다만, 컨테이너 안에 든 파일을 수정하는 것은 애플리케이션에 의도하지 않은 부작용을 초래할 수 있으므로 운영환경에서는 절대 해서는 안된다.
```

- docker container cp
```
COPY: 이미지를 빌드할 때 호스트에서 복사해온 파일을 정의하기 위한 것
docker container cp: 실행 중인 컨테이너와 파일을 주고받기 위한 명령
```

- prune
```
docker container prune
docker image prune
docker system prune
```

#### 4. 운영과 관리를 위한 명령
- docker container stats

#### 5. 도커 컴포즈로 여러 컨테이너 실행하기
- docker-compose 명령으로 컨테이너 실행하기

#### 6. 컴포즈로 여러 컨테이너 실행하기

```
- 젠킨스 컨테이너 실행하기
    volumes항목은 호스트와 컨테이너 사이에 파일을 공유할 수 있는 메커니즘

- 마스터 젠킨스 용 SSH 키 생성
- 슬레이브 젠킨스 컨테이너 생성
```