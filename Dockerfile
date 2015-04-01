FROM scratch

MAINTAINER Chris Moore <chris@cloudspace.com>

ADD ./go-twitter-query /go-twitter-query

CMD chmod +x /go-twitter-query
