package main
// Demonstrate for linkedlist of number

import (
	"fmt"
)

type Number struct {
	body int
	//You can add more element here....
	next *Number
}

type LinkedList struct {
	length	int
	start	*Number
	end		*Number
}

/* Initialize LinkedList*/
func createList() (*LinkedList) {
	return &LinkedList{
		length: 0,
	}
}
/* */

/* Insert as first element */
func (f *LinkedList) insertFirst(newElmt *Number) (){
	if (f.length == 0){
		f.start = newElmt
		f.end = newElmt
	} else {
		currentElmt := f.start
		f.start = newElmt
		f.start.next = currentElmt
	}
	f.length++
}
/* */

/* Insert as last element */
func (f *LinkedList) insertLast(newElmt *Number) (){
	if (f.length == 0){
		f.start = newElmt
		f.end = newElmt
	} else {
		currentElmt := f.end
		currentElmt.next = newElmt
		f.end = currentElmt
	}
	f.length++
}
/* */

/* Insert after idx N */

/* */

/* Get Index */
func (f * LinkedList) getElmt(idx int) (*Number) {
	if (idx >= f.length) {
		return nil
	} else {
		i := 0
		currentElmt := f.start
		for i != idx {
			currentElmt = currentElmt.next
			i++
		}
		return currentElmt
	}
}
/* */

/* Iterate */
func (f *LinkedList) printAllElmt() () {
	currentElmt := f.start
	for currentElmt != nil {
		fmt.Printf("%d ",currentElmt.body)
		currentElmt = currentElmt.next
	}
}
/* */

/* Search Elmt */
func (f *LinkedList) 
/* */

func main() {
	//Testing driver
	l := createList()
	n1 := Number{
		body: 5,
	}
	n2 := Number{
		body: 15,
	}
	n3 := Number{
		body: 20,
	}
	l.insertLast(&n1)
	l.insertLast(&n2)
	l.insertFirst(&n3)

	l.printAllElmt()
	fmt.Println(l.getElmt(2).body)

}