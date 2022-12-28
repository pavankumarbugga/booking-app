# booking-app
A go application

### Initializing a module
`go mod init booking-app`

### Run application
- Run application with one main file main.go
    - `$go run main.go`
- Run application with multiple dependent files main.go depdens on helper.go
    - `$go run main.go helper.go`
    - `$go run .` # Instead of mentioning all files we can run from specifying directory
