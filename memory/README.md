## Memory Key Value Store

### Example

```
import (
	"fmt"

	"github.com/anjulapaulus/kvstore/memory"
)

func main() {
    //memory.WithTTL is optional.
	store := memory.NewMemoryStore("Test", memory.WithTTL(10000000))
	err := store.Set(1, 1)
	if err != nil {
		fmt.Println("Error setting value")
	}
	name := store.Name()
	fmt.Println(name)

	value, err := store.Get(1)
	if err != nil {
		fmt.Println("Error getting value")
	}
	fmt.Println(value)

	cat := store.CreatedAt(1)
	fmt.Println(cat)

	err = store.Delete(1)
	if err != nil {
		fmt.Println("Error deleting value")
	}
}
```

Iterator
```
package main

import (
	"fmt"

	"github.com/anjulapaulus/kvstore/memory"
)

func main() {
	store := memory.NewMemoryStore("Test", memory.WithTTL(10000000))
	store.Set(1, 1)
	store.Set(2, 2)
	store.Set(3, 3)

	i, err := store.NewIterator()
	if err != nil {
		fmt.Println("Error creating iterator")
	}
	// move to next entry
	i.Next()
	fmt.Println(i.Key())   // key of next entry
	fmt.Println(i.Value()) // value of next entry

	// move to previous entry
	i.Previous()
	fmt.Println(i.Key())   // key of previous entry
	fmt.Println(i.Value()) // value of previous entry

	// check if an entry exists next
	check := i.HasNext()
	if check {
		i.Next()
		fmt.Println(i.Key())   // key of next entry
		fmt.Println(i.Value()) // value of next entry
	}

	// check if an entry exists previous
	checkPrevious := i.HasPrevious()
	if checkPrevious {
		i.Previous()
		fmt.Println(i.Key())   // key of next entry
		fmt.Println(i.Value()) // value of next entry
	}

	// get specific key
	i.Seek(1)
	fmt.Println(i.Key())
	fmt.Println(i.Value())

}

```