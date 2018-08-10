FROM golang:alpine
MAINTAINER Raul Perez <repejota@gmail.com>

RUN apk update && \
  apk upgrade && \
  apk add --no-cache git make bash && \
  rm -rf /var/cache/apk/*

ADD . /go/src/github.com/repejota/goomba
WORKDIR /go/src/github.com/repejota/goomba

RUN make deps 
RUN make install 

CMD ["goomba"]
