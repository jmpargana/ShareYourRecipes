FROM golang:latest

RUN apt-get update && apt-get install -y npm
RUN npm i npm@latest -g

RUN mkdir /app

ADD . /app

WORKDIR /app/client

RUN npm i && npm run build

WORKDIR /app/server

RUN go install

CMD ["server"]
