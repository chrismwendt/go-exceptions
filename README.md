# go-exceptions

Ergonomic error handling in Go. Check out [this post](https://dev.to/chrismwendt/ergonomic-error-handling-in-go-using-generics-17b9) for more details.

## Example

Typical error-handling:

```go
func f() error {
  err := g1()
  if err != nil {
    return err
  }

  v1, err := g2()
  if err != nil {
    return err
  }

  // ...
}
```

With go-exceptions:

```go
import ex "github.com/chrismwendt/go-exceptions"

func f() (err error) {
  defer ex.Catch(&err)        // Catch() calls recover() and assigns the error to &err

  ex.Throw(g1())              // Throw() calls panic() if the error argument is not nil
  v1 := ex.Throw1(g2())       // Throw{1,2}() also return values
  v2 := ex.Throw1(g3(), "g3") // Throw() also accepts a label
  // ...
}
```
