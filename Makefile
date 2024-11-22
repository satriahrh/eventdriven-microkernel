KERNEL_ID = kernela
KERNEL_DIR = ./kernels

.PHONY: all clean

all: build

build-kernel-vscode:
	@echo "Building Go plugin..."
	go build -buildmode=plugin -gcflags "all=-N -l" -o $(KERNEL_DIR)/$(KERNEL_ID).so $(KERNEL_DIR)/$(KERNEL_ID)/main.go

build-kernel:
	@echo "Building Go plugin..."
	go build -buildmode=plugin -o $(KERNEL_DIR)/$(KERNEL_ID).so $(KERNEL_DIR)/$(KERNEL_ID)/main.go

clean-kernel:
	@echo "Cleaning up..."
	rm -f $(KERNEL_DIR)/$(KERNEL_ID).so