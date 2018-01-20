APPNAME = blog
BIN = $(GOPATH)/bin
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_UNIX = $(BIN)_unix
COPY = cp
PID = .pid

all: deps build

build:
	rm -rf $(BIN)/$(APPNAME)
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

dev:
	$(GOBUILD) -o $(BIN)/$(APPNAME) -v
	./$(APPNAME)

deps:
	$(GOGET) github.com/tools/godep
	$(BIN)/godep restore

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v


install-ui:
	@yarn install

run:
	$(BIN)/$(APPNAME) & echo $$! > $(PID)

restart:
	@echo restart the app...
	@kill `cat $(PID)` || true
	./$(APPNAME) & echo $$! > $(PID)

kill:
	@kill `cat $(PID)` || true