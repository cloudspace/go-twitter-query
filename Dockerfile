FROM debian:wheezy

MAINTAINER Chris Moore <chris@cloudspace.com>

ADD ./ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD ./microservice.yml /microservice.yml

ADD ./go-twitter-query /go-twitter-query

CMD chmod a+x /go-twitter-query
