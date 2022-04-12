BUILDPATH=$(CURDIR)
BINARY=server

run:
	@echo "run project..."
	go run cmd/server.go

mod:
	@echo "Vendoring..."
	@go mod vendor

build:
	@echo "Compilando..."
	@go build -o $(BUILDPATH)/build/bin/${BINARY} cmd/server.go
	@echo "Binario generado en build/bin/"${BINARY}

clean:
	@echo "Limpiando binarios..."
	@if [ -d $(BUILDPATH)/build/bin ] ; then rm -rf $(BUILDPATH)/build/ ; fi
