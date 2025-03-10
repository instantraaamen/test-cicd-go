# test-cicd-go

```sh
calc-project/
├── main.go
└── operations/
    ├── add.go
    ├── sub.go
    ├── mul.go
    └── div.go
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