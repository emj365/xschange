FROM golang:1.11
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /usr/src/app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . ./
# RUN CGO_ENABLED=0 go build -o main // panic: failed to connect database
RUN go build -o main

FROM alpine:latest
RUN apk add --no-cache libc6-compat
EXPOSE 8000
WORKDIR /root/
COPY --from=0 /usr/src/app/main .
CMD ["./main"]
