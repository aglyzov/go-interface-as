# go-interface-as
A library to convert/parse whatever an interface{} holds into concrete Go types

# examples
```
import "github.com/aglyzov/go-interface-as"
...
as.Bool("True")     -> true,  true
as.Bool("off")      -> false, true
as.Bool("yes ")     -> true,  true
as.Bool(" NO  ")    -> false, true
as.Bool("disabled") -> false, true
as.Bool(153)        -> false, false
as.Bool(1)          -> true,  true
as.Bool(0.0)        -> false, true
...
```
