FROM golang:alpine

WORKDIR /usr/app
ADD ./ ./
RUN go build
EXPOSE 8080
ENTRYPOINT [ "./apiexample" ]
