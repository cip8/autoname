# autoname

Friendly names generator inspired by [moby's](https://github.com/moby/moby/tree/master/pkg/namesgenerator) work.

Function signature:
```
func Generate(delimiter string) string
```

Example usage:
```
import "github.com/cip8/autoname"

fmt.Println(autoname.Generate("-"))
```
