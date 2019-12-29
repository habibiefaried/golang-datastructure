package main
/* 
	Name: Habibie Faried
	Date: 30 December 2019
	Stack implementation
*/

import (
	"fmt"
)

type Element struct {
	body int
	//You can add more element here....
	next *Element
}

type Stack struct {
	length	int
	top		*Element
}

/* Initialize Stack*/
func createStack() (*Stack) {
	return &Stack{
		length: 0,
	}
}
/* */

/*
push(new-item:item-type)
Adds an item onto the stack.
*/
func (f *Stack) push(newElmt *Element) () {
	if (f.length == 0){
		f.top = newElmt
	} else {
		newElmt.next = f.top
		f.top = newElmt
	}
	f.length++
}
/* */

/*
pop()
Removes the most-recently-pushed item from the stack.
*/

func (f *Stack) pop() (*Element) {
	if (f.length == 0){
		return nil
	} else {
		temp := f.top
		f.top = f.top.next
		f.length--
		return temp
	}
}
/* */

/*
getTop():item-type
Returns the last item pushed onto the stack.
*/

func (f *Stack) getTop() (*Element) {
	if (f.length == 0){
		return nil
	} else {
		return f.top
	}
}
/* */

/*
isEmpty():Boolean
True if no more items can be popped and there is no top item.
*/

func (f *Stack) isEmpty() (bool) {
	if (f.top == nil){
		return true
	} else {
		return false
	}
}
/* */

func main(){
	s := createStack()
	i := 0
	var e *Element
	for i != 10 {
		fmt.Println("Pushed ",i*3)
		s.push(&Element{body: i*3})
		i++
	}
	fmt.Println("-------------")

	for !s.isEmpty() {
		e = s.pop()
		fmt.Println("Popped ",e.body)
	}
}