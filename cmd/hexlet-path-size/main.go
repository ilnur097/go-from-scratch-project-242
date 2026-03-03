package main
import (
"fmt"
"os"
"code"
)

func main() {
if len(os.Args) < 2 {
return
}

path := os.Args[1]
result, err := code.GetPathSize(path, false, false, false)
if err != nil {
fmt.Fprintf(os.Stderr, "error: %v\n", err)
os.Exit(1)
}
fmt.Println(result)
}
