package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk traverses t in-order and streams each node's Value on ch, then closes ch.
//
// The inner function walk(n) visits one subtree recursively:
//   - Base case: if n is nil, there is nothing to visit; return without sending.
//   - Otherwise: first walk(n.Left). That call runs to completion before this node is
//     handled, so every value in the left subtree is sent on ch before we proceed.
//     Then send n.Value (the current node's element). Then walk(n.Right), which sends
//     every value in the right subtree in the same in-order pattern.
// Because smaller keys live in the left branch of these BSTs, "left, then node, then right"
// visits keys in ascending order; each tree element is sent exactly once when its node is
// the current n between those two recursive calls.
//
// Run Walk in its own goroutine when comparing trees concurrently (see Same), since sends
// block until a receiver reads.
// https://go.dev/tour/concurrency/7
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(*tree.Tree)
	walk = func(n *tree.Tree) {
		if n == nil {
			return
		}
		walk(n.Left)
		ch <- n.Value
		walk(n.Right)
	}
	walk(t)
}


// Same reports whether t1 and t2 hold the same sequence of values (same contents).
// https://go.dev/tour/concurrency/8
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
}

func printTreeLine(prefix string, isTail bool, n *tree.Tree) {
	if n == nil {
		return
	}
	fmt.Printf("%s", prefix)
	if isTail {
		fmt.Print("└── ")
		prefix += "    "
	} else {
		fmt.Print("├── ")
		prefix += "│   "
	}
	fmt.Println(n.Value)

	if n.Left != nil && n.Right != nil {
		printTreeLine(prefix, false, n.Left)
		printTreeLine(prefix, true, n.Right)
	} else if n.Left != nil {
		printTreeLine(prefix, true, n.Left)
	} else if n.Right != nil {
		printTreeLine(prefix, true, n.Right)
	}
}

func printASCII(root *tree.Tree) {
	if root == nil {
		fmt.Println("(nil)")
		return
	}
	fmt.Println(root.Value)
	prefix := ""
	if root.Left != nil && root.Right != nil {
		printTreeLine(prefix, false, root.Left)
		printTreeLine(prefix, true, root.Right)
	} else if root.Left != nil {
		printTreeLine(prefix, true, root.Left)
	} else if root.Right != nil {
		printTreeLine(prefix, true, root.Right)
	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	fmt.Println("Tree 1 (ASCII):")
	printASCII(t1)
	fmt.Println("\nTree 2 (ASCII):")
	printASCII(t2)
	fmt.Println("\nSame(t1, t2):", Same(t1, t2))

	fmt.Println("\nSame(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))
}
