FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go env -w GOPROXY=https://goproxy.io

RUN go env

RUN go mod download

COPY gomars .

RUN  chmod +x gomars

EXPOSE 8088

CMD ["/app/gomars"]