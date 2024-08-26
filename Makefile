lint:
	golangci-lint run
precommit-hook:
	@[ -f ./precommit-hook.sh ] && cp -v ./precommit-hook.sh ./.git/hooks/pre-commit && echo "precommit hook installed" || echo "error: could not find precommit-hook.sh"

init:
	@echo installing golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3