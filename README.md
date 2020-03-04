# Hello Golang 

### BIN

```bash
go install
``` 

### PKG

```bash
go build path/to/file.go
go install
```

## Testes
FILE_test.go

### Running
```bash
go test
```

## Http server
https://golang.org/pkg/net/http/

```console
import (
    "fmt"
    "net/http"
)
```

### running
```bash
go run server.go
```
>> http://localhost:8080/hello