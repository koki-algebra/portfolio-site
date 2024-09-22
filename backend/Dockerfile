# Development stage
FROM golang:1.22.4 AS dev

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY backend ./

RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.51.0

CMD ["air", "-c", ".air.toml"]


# Production build stage
FROM golang:1.22.4 AS builder

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

# Production stage
FROM alpine:3.19.1 AS deploy

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apk add --no-cache tzdata

COPY --from=builder ${ROOT}/main .

CMD ["/go/src/app/main"]
