// cmd/ciphernexus/main.go
package main

import (
"flag"
"log"
"os"

"ciphernexus/internal/ciphernexus"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ciphernexus.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
