## Tissue for golang - tsh

A bunch of normal tools for golang development

include

* chained error
* multiple error
* concurrence handle
* array helper func
* some file and dir funcs


## Chained error

```golang
import (
    "errors"
    "fmt"
    errs "github.com/yang-zzhong/tsh/errors"
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

## Multiple error

```golang
import (
    "errors"
    "fmt"
    errs "github.com/yang-zzhong/tsh/errors"
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

## Array

```golang
// Contain check whether arr1 contains arr2 without order
arr1 := []int{1, 2, 3}
arr2 := []int{1, 3}
array.Contain(arr1, arr2) // should be true

// Equal check whether arr1 and arr2 equal without order
arr1 := []int{1,2,3}
arr2 := []int{3,2,1}
array.Equal(arr1, arr2) // should be true

// SubFrom get the first index of the arr1 contains arr2 with order, -1 will be returned if not contained

arr1 := []int{1,2,3}
arr2 := []int{2,3}
array.SubFrom(arr1, arr2) // should return 1
```

## concurrence

```golang
arr1 := []int{1,2,3,4,5,6,7,8,9,10}
lock := sync.Mutex
// use 4 goroutine to handle the times
curr.Call(len(arr1), 4, func(start, size int) {
    lock.Lock()
    defer lock.Unlock()
    for i := start; i < start - size; i++ {
        arr1[i] *= arr1[i]
    }
})
```
