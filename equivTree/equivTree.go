package equivTree

import (
	"fmt"
)

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkInt(t, ch)
	close(ch)
}

func WalkInt(t *tree.Tree, ch chan int) {
	if(t==nil){
		return
	}
	WalkInt(t.Left, ch)
	ch <- t.Value
	WalkInt(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	
	for e := range ch1 {
		e2 := <- ch2
		fmt.Println("elems",e,", ",e2)
		if(e!=e2){
			return false
		}
	}
	return true
}
