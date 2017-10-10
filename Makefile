GOOS := linux
ENTRYPOINT := ./cmd/
OUTDIR := ./_build
OUTPUT := $(OUTDIR)/googleDbTest

ifeq ($(OS),Windows_NT)
	OUTPUT := $(OUTPUT).exe
	GOOS := windows
else
  UNAME_S := $(shell uname -s)
  ifeq ($(UNAME_S),Linux)
      GOOS := linux
  endif
  ifeq ($(UNAME_S),Darwin)
      GOOS := darwin
  endif
endif

.PHONY: all clean build test fmt vet

default: all

all: clean fmt vet test build

clean:
	rm -rf $(OUTDIR)
	govendor clean +local

bootstrap:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

build: clean fmt vet
	mkdir -p $(OUTDIR)
	cp cyclotron.toml $(OUTDIR)
	CGO_ENABLED=0 GOOS=$(GOOS) go build -o $(OUTPUT) $(ENTRYPOINT)

test:
	ginkgo -r -randomizeAllSpecs -skipMeasurements=true

fmt:
	govendor fmt +local

vet:
	govendor vet +local