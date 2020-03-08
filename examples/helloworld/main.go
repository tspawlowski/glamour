package main

import (
	"fmt"

	"github.com/tspawlowski/glamour"
)

func main() {
	in := `# Hello World

This is a simple example of glamour!
Check out the [other examples](https://github.com/tspawlowski/glamour/tree/master/examples).

Bye!
`

	out, _ := glamour.Render(in, "dark")
	fmt.Print(out)
}
