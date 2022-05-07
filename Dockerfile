# builder image
FROM golang:1-17-alpine as builder

WORKDIR /app/src/

COPY . .

RUN make build

# runner image
FROM scratch

WORKDIR /app/

COPY --from=builder /main ./main

ENTRYPOINT ["./main"]
