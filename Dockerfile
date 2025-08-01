FROM golang:1.24-alpine
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app .
EXPOSE 8080
CMD ["./app"]
