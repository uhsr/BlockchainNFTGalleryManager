// cmd/blockchainnftgallerymanager/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftgallerymanager/internal/blockchainnftgallerymanager"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftgallerymanager.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
