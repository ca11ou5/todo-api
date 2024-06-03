FROM golang:1.22-alpine as builder
WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "/"]
RUN go mod download

COPY ./ ./
RUN go build -o ./main ./cmd/main.go

FROM alpine:latest as runner

COPY --from=builder /usr/local/src/main /
COPY ./configs/envs /envs

CMD ["/main"]