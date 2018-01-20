APPNAME = blog
BIN = $(GOPATH)/bin
GOCMD = /usr/local/go/bin/go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_UNIX = $(BIN)_unix
COPY = cp

all: deps build

build:
	rm -rf $(bin)/$(APPNAME)
	$(GOBUILD) -o $(BIN)/$(APPNAME) -v
	$(COPY) -R conf/ $(BIN)/conf
	$(COPY) -R views/ $(BIN)/views
	$(COPY) -R static/ $(BIN)/static

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BIN)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BIN)/$(APPNAME) -v
	./$(APPNAME)

deps:
	$(GOGET) github.com/tools/godep
	$(BIN)/godep restore

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v