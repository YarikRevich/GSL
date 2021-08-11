.PHONY: test

stub:
	@echo "It's a stub"
test:
	go test -run tests/*
