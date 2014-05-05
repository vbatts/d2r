
GOPATH := $(shell pwd)/tmp
VERSION := $(shell cat VERSION)
MYPATH := $(shell cat .godir)

all: build

build: .get
	GOPATH=$(GOPATH) go build -ldflags '-X version.VERSION $(VERSION) -X $(MYPATH)/version.VERSION $(VERSION)' .

.get: .gopath
	GOPATH=$(GOPATH) go get -d ./... && touch $@

.gopath: .tmp
	@rm -f $(GOPATH)/src/$(MYPATH) && \
		mkdir -p $(GOPATH)/src/$$(dirname $(MYPATH)) && \
		ln -s $$(pwd) $(GOPATH)/src/$(MYPATH) && \
		echo export GOPATH=$(GOPATH) && \
		touch $@


.tmp:
	@mkdir -p ./tmp && touch $@

clean:
	rm -rf ./tmp .get .tmp .gopath

