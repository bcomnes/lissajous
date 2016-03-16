// lisssajous generator from "The Go programming langauge" book
package main

import (
	"github.com/bcomnes/lissajous"
	"os"
)

func main() {
	lisssajous.Draw(os.Stdout)
}
