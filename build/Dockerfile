FROM ubuntu

MAINTAINER Andrey Chernih <andrey.chernih@gmail.com>

RUN apt-get update && apt-get install -y ca-certificates
RUN mkdir /gouscis

WORKDIR /gouscis

ADD . /gouscis/

CMD ["/gouscis/gouscis.linux"]
