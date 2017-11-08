package main

import (
	"fmt"
	"github.com/RubenScholten/Training/04_scope/01_package-scope/02_visibility/vis"
)

func main() {

	fmt.Println(vis.Myname)
	vis.PrintVar()
}
