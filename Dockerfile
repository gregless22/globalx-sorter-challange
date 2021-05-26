FROM golang:latest AS builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/main.go

FROM scratch
COPY --from=builder /build/main /app/
COPY unsorted-names-list.txt /app/unsorted-names-list.txt
WORKDIR /app
EXPOSE 9010
CMD ["./main", "./unsorted-names-list.txt"]
