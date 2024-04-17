FORMAT ?= csv
CLUSTER_ENV ?= baremetal
DEST_DIR ?= .
DEPLOYMENT ?= mno
GO_SRC := cmd/main.go

EXECUTABLE := commatrix-gen

.DEFAULT_GOAL := run

build:
	go build -o $(EXECUTABLE) $(GO_SRC) 

# TODO: check if oc is installed
generate: build
	mkdir -p $(DEST_DIR)/communication-matrix
	./$(EXECUTABLE) -format=$(FORMAT) -env=$(CLUSTER_ENV) -destDir=$(DEST_DIR)/communication-matrix -deployment=$(DEPLOYMENT)

clean:
	@rm -f $(EXECUTABLE)
