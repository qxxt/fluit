# fluit
fluit is a simple regex based text formatter for wrapping long text, adding margin and creating cli usage for Golang.

### Basic usage

```
package main

import (
	"fmt"
	"github.com/qxxt/fluit"
)

func main() {
	fluit.SetBreakpoint(70)
	fmt.Print(fluit.Sprintwrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim. Donec malesuada accumsan porttitor. In neque libero, sollicitudin vitae interdum ut, semper lacinia mauris. Integer auctor, nisl at commodo feugiat, justo urna varius arcu, a lacinia mi mauris sit amet purus."))

	fmt.Print("\nUsages:\n")
	u := fluit.Usgs{}
	u.AddUsg("lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("sit", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.AddUsg("amet", "Donec malesuada accumsan porttitor.")
	u.AddUsg("consectetur", "Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.PrintUsg()

	fmt.Print("\nAditional flags:\n")
	u.SetArgLen(20)
	u.AddUsg("--lorem string", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("--ipsum int", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	u.AddUsg("--sit=sits", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.AddUsg("--amet, -a", "Donec malesuada accumsan porttitor.")
	u.AddUsg("--consectetur", "Nullam nisi magna, fermentum sit amet quam id, scelerisque elementum enim.")
	u.PrintUsg()
}
```

This will output the following to console:

```
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam
    nisi magna, fermentum sit amet quam id, scelerisque elementum
    enim. Donec malesuada accumsan porttitor. In neque libero,
    sollicitudin vitae interdum ut, semper lacinia mauris. Integer
    auctor, nisl at commodo feugiat, justo urna varius arcu, a lacinia
    mi mauris sit amet purus.

Usages:
  lorem        Lorem ipsum dolor sit amet, consectetur adipiscing
               elit.
  ipsum        Lorem ipsum dolor sit amet, consectetur adipiscing
               elit.
  sit          Lorem ipsum dolor sit amet, consectetur adipiscing
               elit. Nullam nisi magna, fermentum sit amet quam id,
               scelerisque elementum enim.
  amet         Donec malesuada accumsan porttitor.
  consectetur  Nullam nisi magna, fermentum sit amet quam id,
               scelerisque elementum enim.

Aditional flags:
  --lorem string        Lorem ipsum dolor sit amet, consectetur
                        adipiscing elit.
  --ipsum int           Lorem ipsum dolor sit amet, consectetur
                        adipiscing elit.
  --sit=sits            Lorem ipsum dolor sit amet, consectetur
                        adipiscing elit. Nullam nisi magna, fermentum
                        sit amet quam id, scelerisque elementum enim.
  --amet, -a            Donec malesuada accumsan porttitor.
  --consectetur         Nullam nisi magna, fermentum sit amet quam id,
                        scelerisque elementum enim.
```

For more about documentation, please go to [hh](https://pkg.go.dev/github.com/qxxt/fluit)