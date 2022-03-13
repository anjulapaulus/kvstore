[![codecov](https://codecov.io/gh/anjulapaulus/kvstore/branch/main/graph/badge.svg?token=zXnY4isAr9)](https://codecov.io/gh/anjulapaulus/kvstore)

# KV Store
The project holds key-value store implementations.
1. Memory key-value store

### Examples

```
package main

import "github.com/anjulapaulus/kvstore/memory"

func main(){
    store := memory.NewMemoryStore("Test")
}
```
