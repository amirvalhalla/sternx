ifneq (,$(filter $(OS),Windows_NT MINGW64))
EXE = .exe
RM = del /q
else
RM = rm -rf
endif

devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.12
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.12
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1
	go install github.com/bufbuild/buf/cmd/buf@v1.25.0
	go install mvdan.cc/gofumpt@latest
	go install github.com/rakyll/statik@v0.1


### proto
proto:
	$(RM) www/grpc/gen
	cd www/grpc/buf && buf generate --template buf.gen.yaml ../proto
	cd www/grpc/ && statik -m -f -src swagger-ui/

# Generate static assets for Swagger-UI
	cd www/grpc/ && statik -m -f -src swagger-ui/

fmt:
	gofumpt -l -w .

check:
	golangci-lint run --build-tags "${BUILD_TAG}" --timeout=20m0s
