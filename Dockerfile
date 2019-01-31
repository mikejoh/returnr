# Build
FROM golang:alpine AS builder
ADD main.go /src
RUN cd /src && go build -o returnr

# Final
FROM alpine
WORKDIR /app
COPY --from=builder /src/returnr /app/
ENTRYPOINT ["./returnr"]
