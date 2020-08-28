# Compile stage
FROM golang:1.13.8 AS build-env
# ADD . /dockerdev
COPY . /usr/local/go/src/gosharedlib/
WORKDIR /usr/local/go/src/gosharedlib/
# CMD ["cp", "test.so", "/test.so"]
RUN go build -o gosharedlib.so -buildmode=c-shared