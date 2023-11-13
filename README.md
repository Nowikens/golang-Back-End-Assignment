# General information
1. I threw away macOS related files, as I did this on Ubuntu.
2. Package tries to process as much as it can, if it can't handle corrupted email or row it'll go to the next row. The reason is that example data contains column names across document, and it would be wastefull to process document only to the point of first corrupted email.

# Run application
```bash
go run cmd/main.go -file=<path_to_file>
```
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
    Fasteting parsing csv data by using goroutines - which I didn't manage to achieve unfortunatelly.
