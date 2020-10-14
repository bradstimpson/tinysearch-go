FROM golang:alpine as builder
RUN mkdir /build 
WORKDIR /build
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh bash
RUN git clone https://github.com/bradstimpson/tinysearch-go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s -extldflags "-static"' -o main .
FROM debian 
COPY --from=builder /build/* /app/
WORKDIR /app
EXPOSE 5001
CMD tail -f /dev/null
# CMD ["./main"]