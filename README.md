Fully compatable with github.com/pkg/errors@v0.9.1 5dd12d0


## New Features

1. `SetSkipCallers`: set global depth to skip callers
2. `WrapWithSkip`:
3. `WrapfWithSkip`:


### Set global depth to skip callers

```go
import "github.com/Laisky/errors"

errors.SetSkipCallers(3)  // default to 3
```
