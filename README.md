# istor

`istor` is a tool to check if an IP address is a Tor exit node.

## CLI usage

### Installation

Download latest binaries from https://github.com/igolaizola/istor/releases

or install using go

```
go install github.com/igolaizola/istor/cmd/istor@latest
```

### Usage

```
$ istor --ip 11.22.33.44
false
```

## Library usage

```golang
import (
    "context"
	
    "github.com/igolaizola/istor"
)

func main() {
    ip := "11.22.33.44"
    err, isTor := istor.IsExitNode(context.Background(), ip)
    if err != nil {
        log.Fatal(err)
    }
    if isTor {
        fmt.Printf("%s is a Tor exit node\n", ip)
    } else {
        fmt.Printf("%s is not a Tor exit node\n", ip)
    }
}
```
