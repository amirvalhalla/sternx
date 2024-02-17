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
- You can change GRPC Gateway port , DSN etc. at `confg.json` file in root of the project.
- Tip: if you changed the GRPC Gateway port or GRPC Server port, please read Contribution Guide
or you need to do 2 steps:
```bash
make devtools
make proto
```

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
- For building the proto files Run below command
```bash
make proto
```
