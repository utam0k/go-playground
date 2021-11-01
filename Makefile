GO ?= go

GO_BUILD := $(GO) build

.PHONY: mycat typing

mycat:
	$(GO_BUILD) -o mycat ./cmd/mycat
	chmod +x mycat

typing:
	$(GO_BUILD) -o typing ./cmd/typing
	chmod +x typing

clean:
	rm mycat typing