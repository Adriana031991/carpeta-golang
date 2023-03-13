package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Sistema operativo: %v\nArquitectura: %v\n", runtime.GOOS, runtime.GOARCH)
}
