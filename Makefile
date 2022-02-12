# jotto
.POSIX:

include config.mk

all: clean build

build:
	go build -ldflags "-X main.Version=$(VERSION)" $(GOFLAGS)
	scdoc < jotto.6.scd | sed "s/VERSION/$(VERSION)/g" > jotto.6

clean:
	rm -f jotto
	rm -f jotto.6

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f jotto $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/jotto
	mkdir -p $(DESTDIR)$(MANPREFIX)/man6
	cp -f jotto.6 $(DESTDIR)$(MANPREFIX)/man6/jotto.6
	chmod 644 $(DESTDIR)$(MANPREFIX)/man6/jotto.6

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/jotto
	rm -f $(DESTDIR)$(MANPREFIX)/man6/jotto.6

.PHONY: all build clean install uninstall
