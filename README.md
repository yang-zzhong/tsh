## Tissue for golang - tsh

A bunch of normal tools for golang development

include

* chained error
* multiple error
* concurrence handle
* array helper func
* some file and dir funcs


## chained error

```golang
import (
    "errors"
    "fmt"
    errs "git.woa.com/oliverzyang/tsh/errors"
)

func LackOfBlock() bool {
    return true
}

func BuildWorld() error {
    if LackOfBlock() {
        return errors.New("lack of block")
    }
    return nil
}

func SayHello() errs.ChainedError {
    if err := BuildWorld(); err != nil {
        return errs.New(err).Cause("build world failed")
    }
    fmt.Printf("hello world!!\n")
    return nil
}

func main() {
    if err := SayHello(); err != nil {
        err.Cause("say hello failed").Print()
        return
    }
}
```

## multiple error

```golang
import (
    "errors"
    "fmt"
    errs "git.woa.com/oliverzyang/tsh/errors"
)

func SayHello() errs.MultiError {
    err := errs.Multiple()
    if !Build1World() {
       err.Occurred(errors.New("build 1st world error")) 
    }
    if !Build2World() {
       err.Occurred(errors.New("build second world error")) 
    }
    if err.AnyOccurred() {
        return err
    }
    fmt.Printf("hello world!!\n")
}

func main() {
    if err := SayHello(); err != nil {
        fmt.Printf("say hello error: %s\n", err.Error())
    }
}
```
