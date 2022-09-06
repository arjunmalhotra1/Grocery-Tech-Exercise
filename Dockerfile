FROM golang:1.15-alpine3.14

WORKDIR /usr/src/Grocery-TECH-EXERCISE
ENV GO111MODULE=on
COPY ./ ./
RUN apk add bash ca-certificates git gcc g++ libc-dev make
RUN go mod download
CMD ["/bin/bash"]