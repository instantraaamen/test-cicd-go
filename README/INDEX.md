# 0. repository tree
```sh
.
├── Dockerfile
├── README_DOCKER.md
├── README_FIRST.md
├── README_SECOND.md
├── go.mod
├── main.go
└── operations
    ├── add.go
    ├── add_test.go
    ├── div.go
    ├── div_test.go
    ├── mul.go
    └── sub.go
```

# 1. module init
```sh
go mod init test-cicd-go
```

# 2. fix import project
```go
// main.go
import (
    "fmt"
    "test-cicd-go/operations"
)
...
```

# 3. exaction program
```sh
go run main.go
```