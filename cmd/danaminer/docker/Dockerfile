# -- multistage docker build: stage #1: build stage
FROM golang:1.23-alpine AS build

RUN mkdir -p /go/src/github.com/danannet/danad

WORKDIR /go/src/github.com/danannet/danad

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/danannet/danad/cmd/danaminer
RUN GOOS=linux go build -a -installsuffix cgo -o danaminer .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/danannet/danad/cmd/danaminer/danaminer /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
