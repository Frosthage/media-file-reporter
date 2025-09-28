FROM golang:1.23

RUN apt-get update && apt-get install -y \
    gcc-mingw-w64 \
    && apt-get clean

WORKDIR /app

COPY . .

ENV GOOS=windows \
    GOARCH=amd64 \
    CGO_ENABLED=1 \
    CC=x86_64-w64-mingw32-gcc
