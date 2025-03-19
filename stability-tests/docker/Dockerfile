ARG DANAD_IMAGE
ARG DANAMINER_IMAGE

FROM ${DANAD_IMAGE} as danad
FROM ${DANAMINER_IMAGE} as danaminer

FROM golang:1.23-alpine

RUN mkdir -p /go/src/github.com/danannet/danad

WORKDIR /go/src/github.com/danannet/danad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=danad /app/ /app/
COPY --from=danaminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/danannet/danad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]