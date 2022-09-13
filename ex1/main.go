package main

import (
	"fmt"
)

import "golang.org/x/tour/tree"
import "github.com/damientl/golangEx/equivTree"

func main() {
	
	passed := equivTree.Same(tree.New(1), tree.New(1)) && ! equivTree.Same(tree.New(1), tree.New(2)) 
	
	fmt.Println(passed)
	
}
