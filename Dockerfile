FROM golang:1-alpine

WORKDIR /app
COPY . .

RUN go build -o jaytube-content-information-service cmd/jaytube_content_information_service/main.go

CMD ["./jaytube-content-information-service"]