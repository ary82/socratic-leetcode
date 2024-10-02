FROM golang:alpine3.20 AS base

WORKDIR /app

COPY . .

RUN apk add make

RUN make clean
RUN go mod download
RUN make build

FROM alpine:3.20
COPY --from=base /app/main /app/

WORKDIR /app

CMD [ "./main" ]
