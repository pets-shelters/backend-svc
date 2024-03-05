FROM golang:1.19
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app -o /docker-gs-ping
CMD ["/docker-gs-ping"]