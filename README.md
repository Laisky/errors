# Errors with stack info

| Version | Support Go |
| ------- | ---------- |
| v1      | >= 1.13    |
| v2      | >= 1.20    |

## New Features

### v2

1. `Join`
2. `Is`

### v1

1. `SetSkipCallers`: set global depth to skip callers
2. `WrapWithSkip`:
3. `WrapfWithSkip`:


### Set global depth to skip callers

```go
import "github.com/Laisky/errors/v2"

errors.SetSkipCallers(3)  // default to 3
```
