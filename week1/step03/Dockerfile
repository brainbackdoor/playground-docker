FROM node:carbon

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# 앱 의존성 설치
COPY package*.json ./
RUN npm install

# 앱 소스 추가
COPY index.js .

EXPOSE 8080
CMD [ "npm", "start" ]
