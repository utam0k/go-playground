GO ?= go

GO_BUILD := $(GO) build

.PHONY: mycat

mycat:
	$(GO_BUILD) -o mycat .

clean:
	rm mycat