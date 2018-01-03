GO=go
GOFLAGS=-gcflags "-N -l"
all: spider_server
mac: spider_server_mac

.PHONY: all test clean proto

SRC=$(shell find SpiderServer -maxdepth 1 -name "*.go" -type f)

spider_server: $(SRC) spider_server.go
	$(GO) version
	start=$$(date +%s);\
	env GOOS=linux GOARCH=amd64 $(GO) build -o $@ $(GOFLAGS);\
	end=$$(date +%s);\
	echo "time usage:"$$(( $$end - $$start))"s"


spider_server_mac : $(SRC) spider_server.go
	$(GO) build -o $@ $(GOFLAGS);\

proto:
	mkdir -p $(GOPATH)/src/spider_data.pb
	(cd proto; protoc --go_out=$(GOPATH)/src/spider_data.pb spider_data.proto)

dep:
	$(GO) get -u github.com/golang/protobuf/proto
	$(GO) get -u github.com/julienschmidt/httprouter
	$(GO) get -u github.com/urfave/negroni
	$(GO) get -u github.com/PuerkitoBio/goquery
	$(GO) get -u github.com/gogo/protobuf/proto
	$(GO) get -u github.com/headzoo/surf
	$(GO) get -u github.com/headzoo/surf/browser

clean:
	rm  ./spider_server

install_golang_mac:
	echo 'Updating homebrew...'
	brew update
	echo 'Installing protobuf..'
	brew install protobuf
	echo 'Installing golang...'
	brew install golang

	echo 'Add environment variables...'
	echo 'export GOPATH="$$HOME/go-workspace"' >> ~/.profile
	echo 'export GOROOT="/usr/local/opt/go/libexec"' >> ~/.profile
	echo 'export PATH="$$PATH:$$GOPATH/bin"' >> ~/.profile
	echo 'export PATH="$$PATH:$$GOROOT/bin"' >> ~/.profile
	source ~/.profile

	echo 'Clear & Create go workspace...'
	mkdir -p $$GOPATH $$GOPATH/src $$GOPATH/pkg $$GOPATH/bin
	rm -rf $$GOPATH/*
	cp -R * $$GOPATH
	echo 'Finished!!!'
