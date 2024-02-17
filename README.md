# Sternx

[![Lint and format check](https://github.com/amirvalhalla/sternx/actions/workflows/linting.yml/badge.svg?branch=main)](https://github.com/amirvalhalla/sternx/actions/workflows/linting.yml)

### Requirements
- go >= 1.21.4
- Makefile (for development)

### Usage

- run below command to start GRPC Gateway and GRPC Server
```bash
go run cmd/main.go
```
- By default, GRPC Gateway accessible at `localhost:8080` and GRPC Server accessible at
`loclahost:50051`

### Contribution Guide
- Install Make tool
- Run below command to install dev tools
```bash
make devtools
```
- for formatting the code Run
```bash
make fmt
```
- for linting check Run
```bash
make check
```
