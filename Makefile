# Makefile


# Run tests with coverage and generate HTML report
coverage:
	go test ./... -v -coverpkg=./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Clean up coverage files
clean:
	rm -f coverage.out coverage.html