FROM golang:1.24-alpine AS build

WORKDIR app

COPY . ./

RUN go build -o main main.go

FROM alpine:3.22 AS run

WORKDIR /app

COPY --from=build /go/app/main .
COPY --from=build /go/app/schema.json .

RUN chmod +x main

ENTRYPOINT ["/app/main"]