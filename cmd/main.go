package main

import (
	"fmt"

	ct "github.com/vadervela/compiler-test/types"
	tt "github.com/vadervela/types-test/types"
)

func main() {
	_ = tt.Test{}
	_ = ct.Test{}

	fmt.Println("hello from server: depends on types and compiler")
}
