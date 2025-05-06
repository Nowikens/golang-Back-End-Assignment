# General information
1. I threw away macOS related files, as I did this on Ubuntu.
2. Package tries to process as much as it can, if it can't handle corrupted email or row it'll go to the next row. The reason is that example data contains column names across document, and it would be wastefull to process document only to the point of first corrupted email.

# Usage examples:
1. output to console
```go run cmd/main.go --file=./customers.csv```
2. output to file
```go run cmd/main.go --file=./customers.csv --output=./test.json```
3. format flag example (currently only json works)
```go run cmd/main.go --file=./customers.csv --output=./test.json --format=json```
# Tests
```bash
go test ./...
```
# Benchmarks
```bash
go test ./... -bench=. -benchmem
```

# Used resources
1. https://mecitsemerci.net/how-to-create-a-csv-file-in-memory-with-golang
    Creating small csv in memory
2. https://golangbyexample.com/sort-custom-struct-collection-golang/
    Sorting custom structs
3. https://medium.com/@snassr/processing-large-files-in-go-golang-6ea87effbfe2
    Fastening parsing csv data by using goroutines - which I didn't manage to achieve unfortunatelly.
4. (*NEW 2025*) https://boyter.org/posts/golang-slog-disable-tests/ https://x.com/ohmypy/status/1866078171340185706 
    Turn off logging while running tests


# Changes 2025
1. Updated go version to 1.24
2. Replaced bufio with encoding/csv
3. Allowed any email column position - this change also made headers line required
4. New CLI flags
    - --format - previously it was just printing go's slice of structs, now it's json. JSON is now the only supported and default format but code was written with extending to other formats in mind.
    - --output - previously output was printed to console, now can be written to file
5. Tests
    - increased code coverage
    - some minor changes in existing tests to comply with code changes
6. Created Makefile
    - one command coverage html report generation
    ```make coverage```
7. Supressed logger output to console while running tests