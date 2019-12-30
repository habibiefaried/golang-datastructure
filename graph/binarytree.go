package main
/* 
	Name: Habibie Faried
	Date: 30 December 2019
	Binary Tree implementation. You can use this to sort array of int as well
*/

import (
	"fmt"
	"math/rand"
)

type Node struct {
	body int
	//You can add more element here....
	left *Node
	right *Node
}

/*
printTree()
Printing the tree node, and calling branch using recursive
*/
func (f * Node) printTree() (){
	if f.left != nil {
		f.left.printTree()
	}
	
	fmt.Printf("%d ",f.body)

	if f.right != nil {
		f.right.printTree()
	}
}

/*
insert()
binary insert, Compare the new value with the parent node
*/

func (f * Node) insert(e * Node) () {
	if f.body > e.body {
		if (f.left == nil) {
			f.left = e
		} else {
			f.left.insert(e)
		}
	} else {
		if (f.right == nil) {
			f.right = e
		} else {
			f.right.insert(e)
		}
	} //Ignore if equal
}

func main(){
	root := Node{body: 15}
	i := 0
	var t int
	fmt.Println("=======================")
	fmt.Println("   RANDOM LIST NUMBER  ")
	fmt.Println("=======================")
	for i < 100 {
		t = rand.Intn(1000)
		fmt.Printf("%d ",t)
		root.insert(&Node{body: t})
		i++
	}
	fmt.Println()
	fmt.Println("=======================")
	fmt.Println("       TREE PRINT      ")
	fmt.Println("=======================")
	root.printTree()
}