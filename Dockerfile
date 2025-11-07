# build stage
FROM golang:1.25 AS builder

WORKDIR /app

RUN git clone https://github.com/golang-migrate/migrate.git /app/migrate
RUN cd /app/migrate/cmd/migrate && go build -tags 'postgres' -o /go/bin/migrate .

# final stage
FROM golang:1.25

WORKDIR /app

RUN apt-get update && apt-get install -y postgresql-client

RUN go install github.com/air-verse/air@latest

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

COPY . .
RUN go mod tidy