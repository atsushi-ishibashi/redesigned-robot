FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/github.com/atsushi-ishibashi/redesigned-robot
RUN \
  apk update && \
  rm -rf /var/cache/apk/* && \
  mkdir -p $APPDIR
ADD . $APPDIR/
WORKDIR $APPDIR
RUN go build -ldflags "-s -w" -o redesigned-robot .

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/atsushi-ishibashi/redesigned-robot/redesigned-robot ./
ENTRYPOINT ["./redesigned-robot"]
