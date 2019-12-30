package main
/* 
	Name: Habibie Faried
	Date: 30 December 2019
	Simple doubly linkedlist (next and prev) with double head-tail pointer
*/

import (
	"fmt"
)

type Element struct {
	body int
	//You can add more element here....
	next *Element
	prev *Element
}

type LinkedList struct {
	length	int
	start	*Element
	end		*Element
}

/* Initialize LinkedList*/
func createList() (*LinkedList) {
	return &LinkedList{
		length: 0,
	}
}
/* */

/* Insert as first element */
func (f *LinkedList) insertFirst(newElmt *Element) () {
	if (f.length == 0){
		f.start = newElmt
		f.end = newElmt
	} else {
		currentElmt := f.start
		currentElmt.prev = newElmt

		f.start = newElmt
		f.start.next = currentElmt
	}
	f.length++
}
/* */

/* Insert as last element */
func (f *LinkedList) insertLast(newElmt *Element) () {
	if (f.length == 0){
		f.start = newElmt
		f.end = newElmt
	} else {
		currentElmt := f.end
		currentElmt.next = newElmt
		newElmt.prev = f.end
		f.end = newElmt
	}
	f.length++
}
/* */

/* Insert after index N */
func (f * LinkedList) insertAfterIdx(idx int, newElmt *Element) () {
	if (idx < f.length) {
		if (idx == f.length - 1){
			f.insertLast(newElmt)
		} else {
			el := f.getElmt(idx)
			newElmt.prev = el
			newElmt.next = el.next
			el.next.prev = newElmt
			el.next = newElmt
			f.length++
		}
	}
}
/* */

/* Get Index */
func (f * LinkedList) getElmt(idx int) (*Element) {
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

/* Delete First */
func (f * LinkedList) delFirst() (*Element) {
	// will return deleted element
	if (f.length == 0) {
		return nil
	} else {
		currentElmt := f.start
		f.start = f.start.next
		f.start.prev = nil //Remove previous pointer
		f.length--
		return currentElmt
	}
}
/* */

/* Delete Last */
func (f * LinkedList) delLast() (*Element) {
	// will return deleted element
	if (f.length == 0){
		return nil
	} else {
		tmp := f.end
		currentElmt := f.end.prev
		f.end = currentElmt
		f.end.next = nil
		f.length--
		return tmp
	}
}
/* */

/* Delete index N */
func (f * LinkedList) delIdx(idx int) (*Element){
	// will return deleted element
	if (f.length == 0){
		return nil
	} else if (idx < f.length) {
		if (idx == f.length - 1){
			return f.delLast()
		} else if (idx == 0){
			return f.delFirst()
		} else {
			el := f.getElmt(idx)
			el.prev.next = el.next
			el.next.prev = el.prev
			f.length--
			return el
		}
	} else {
		return nil
	}
}
/* */

/* Iterate forward with next pointer */
func (f *LinkedList) printAllElmtForward() () {
	if (f.length == 0) {
		fmt.Printf("List empty")
	} else {
		fmt.Printf("Forward (%d): ",f.length)
		currentElmt := f.start
		for currentElmt != nil {
			fmt.Printf("%d ",currentElmt.body)
			currentElmt = currentElmt.next
		}
		fmt.Println()
	}
}
/* */

/* Iterate Backward using prev pointer */
func (f *LinkedList) printAllElmtBackward() () {
	if (f.length == 0) {
		fmt.Printf("List empty")
	} else {
		fmt.Printf("Backward (%d): ",f.length)
		currentElmt := f.end
		for currentElmt != nil {
			fmt.Printf("%d ",currentElmt.body)
			currentElmt = currentElmt.prev
		}
		fmt.Println()
	}
}
/* */

/* is Element Exists */
func (f *LinkedList) isElmtExists(e *Element) (bool){
	currentElmt := f.start
	for currentElmt != nil {
		if (currentElmt.body == e.body){
			return true
		}

		currentElmt = currentElmt.next
	}
	return false
}
/* */

func main() {
	//Only for testing
	l := createList()
	i := 0
	for i != 10 {
		l.insertFirst(&Element{body: i*2})
		i++
	}

	i = 0
	for i != 10 {
		l.insertLast(&Element{body: i*-1})
		i++
	}

	nsearch := Element{body: -1}

	l.printAllElmtForward()
	l.printAllElmtBackward()

	fmt.Println("Getting idx 2: ",l.getElmt(2).body)
	fmt.Println(l.isElmtExists(&nsearch))

	e := l.delFirst()
	fmt.Println("Deleted elmt ",e.body)
	e = l.delFirst()
	fmt.Println("Deleted elmt ",e.body)

	l.printAllElmtForward()
	l.printAllElmtBackward()

	e = l.delLast()
	fmt.Println("Deleted elmt ",e.body)
	e = l.delLast()
	fmt.Println("Deleted elmt ",e.body)
	e = l.delLast()
	fmt.Println("Deleted elmt ",e.body)

	l.printAllElmtForward()
	l.printAllElmtBackward()

	fmt.Println("insertAfterIdx")
	l.insertAfterIdx(0,&Element{body: 200})
	l.insertAfterIdx(2,&Element{body: 201})

	l.printAllElmtForward()
	l.printAllElmtBackward()

	e = l.delIdx(1)
	fmt.Println("Deleted elmt ",e.body)
	e = l.delIdx(2)
	fmt.Println("Deleted elmt ",e.body)
	e = l.delIdx(3)
	fmt.Println("Deleted elmt ",e.body)
	l.printAllElmtForward()
	l.printAllElmtBackward()
}