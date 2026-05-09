```Plug In```
It is little piece of code that pull in at runtime after the main stars

```Packages```
Packages can not have cyclic dependencies, and anything inside 
package that starts with Capital letter is exported and small case one are private

```Declaration Block```

var (
	x,y int
	z float

	..so on
)


```GO keeps arbitrary precision (very high)```

``` Reading input``` 
just single element fmt.Scan(&x)


entire line 


package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	if(sc.Text()==""){
		sc.Scan()
	}
	return sc.Text()
}

func main() {
	line := nextLine()
	fmt.Println("Result:", line)
}
