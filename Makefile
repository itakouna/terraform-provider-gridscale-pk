PROJECT=terraform-provider-gridscale
PROVIDER=gridscale
ARCH=amd64
VERSION=0.1.0
EXTENSION=""

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
PROJECT_DIR=${GOPATH}/src/github.com/${PROVIDER}/${PROJECT}/
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

default: build

ifeq ($(OS),Windows_NT)
EXTENSION=".exec"
endif

build:
	    GOGC=off CGOENABLED=0 godep go build -i -o $(PROJECT)$(EXTENSION)

release:
	cd ${PROJECT_DIR}; \
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${PROJECT}-linux-${ARCH} .

	cd ${PROJECT_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${PROJECT}-darwin-${ARCH} .

	cd ${PROJECT_DIR}; \
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${PROJECT}-windows-${ARCH}.exe .

copy:
	    cp $(PROJECT_DIR)$(PROJECT)$(EXTENSION) $(GOPATH)/bin

install: build copy

clean:
	-rm -f ${PROJECT}-*

.PHONY: build release install
