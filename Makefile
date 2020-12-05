GOBUILD ?= go build

all: main.go logging.go keyExchange.go keyCrypto.go files.go fileCrypto.go deviceDiscovery.go config.go
	$(GOBUILD)

install: opensend
	install -Dm755 opensend $(DESTDIR)/usr/bin/opensend

install-macos: opensend
	install -m755 opensend $(DESTDIR)/usr/local/bin/opensend