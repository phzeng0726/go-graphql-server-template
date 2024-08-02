# Makefile

.PHONY: help graphql-generate

help:
	@echo "Available targets:"
	@echo "  graphql-generate   - Run 'go get' and then 'go run' for gqlgen"
	@echo "Using `make graphql-generate`"

graphql-generate:
	go get github.com/99designs/gqlgen@v0.17.49
	go run github.com/99designs/gqlgen