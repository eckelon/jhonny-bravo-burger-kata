.PHONY: test
test:
	go test -v -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
clean:
	rm *.out
