package main

import (
	"os"

	"github.com/taylormonacelli/allit"
)

func main() {
	code := allit.Execute()
	os.Exit(code)
}
