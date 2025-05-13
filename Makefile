# Makefile


# Run tests with coverage and generate HTML report
coverage:
	go test ./... -v -coverpkg=./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html



# Clean up coverage files
clean:
	rm -f coverage.out coverage.html


# Run benchmarks and save output to a specified file
# e.g make benchmark file=after_changes
benchmark:
	go test ./... -bench=. -benchmem > benchmarks/$(file)\.txt

# comapring 2 benchmark files
# e.g. make benchmark_comp file1=pre_changes file2=after_changes
benchmark_comp:
	mkdir -p benchmarks
	benchstat benchmarks/$(file1)\.txt  benchmarks/$(file2)\.txt > benchmarks/$(file1)__$(file2)_comparison\.txt