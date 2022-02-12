# jotto
.POSIX:

include config.mk

all: clean build

build:
	go build

clean:
	rm -f jotto

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f jotto $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/jotto

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/jotto

.PHONY: all build clean install uninstall
