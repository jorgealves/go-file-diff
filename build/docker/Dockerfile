FROM golang:buster
ENV APP_SOURCE=/go/src/app
WORKDIR $APP_SOURCE
ADD . $APP_SOURCE
RUN go mod verify
RUN go test -v
RUN go build -v
RUN go install -v

CMD ['go-file-diff']
