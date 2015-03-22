DESTDIR ?= /usr/local

all: build

build: gouscis

gouscis: *.go
				 go fmt
				 go build

clean:
				 rm -f gouscis

install: all
				 install -d $(DESTDIR)/bin
				 install gouscis $(DESTDIR)/bin
