# 0. repository tree for target
```sh
operations/
├── add.go
├── add_test.go
├── sub.go
├── sub_test.go
├── mul.go
├── mul_test.go
├── div.go
└── div_test.go
```

# 1. exaction program
```sh
# partial
go test ./operations -v
# all
go test ./... -v
```

# 2.  confirm coverage & report
```sh
# include coverage test
go test -coverprofile=coverage.out ./...

# create html from coverage report
go tool cover -html=coverage.out -o coverage.html
```