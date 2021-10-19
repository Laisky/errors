Fully compatable with github.com/pkg/errors@v0.9.1 5dd12d0


## New Features

0. set depth to skip callers


### Set depth to skip callers

```go
import "github.com/Laisky/errors"

errors.SetSkipCallers(3)  // default to 3
```
