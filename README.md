# github.com/mitchallen/go-lib

## Usage

### Initialize your module

```sh
go mod init example.com/my-golib-demo
```

### Get the go-lib module

Note that you need to include the **v** in the version tag.

```sh
go get github.com/mitchallen/go-lib@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/mitchallen/go-lib"
)

func main() {
    fmt.Println(golib.Add(2,3))
    fmt.Println(golib.Subtract(2,3))
}
```

## Testing

```sh
go test
```

## Tagging

```sh
git tag v0.1.0
git push origin --tags
```# go-lib-test
# ecommerce_go_common
