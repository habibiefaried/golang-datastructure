package main
/* 
	Name: Habibie Faried
	Date: 30 December 2019
	Queue implementation
*/

import (
	"fmt"
)

type Element struct {
	body int
	//You can add more element here....
	next *Element
}

type Queue struct {
	length	int
	head	*Element
	tail	*Element
}

/* Initialize Queue*/
func createQueue() (*Queue) {
	return &Queue{
		length: 0,
	}
}
/* */

/* 
enqueue
Adds an item onto the end of the queue.
*/
func (f *Queue) enqueue(newElmt *Element) () {
	if (f.length == 0){
		f.head = newElmt
		f.tail = newElmt
	} else {
		lastElmt := f.tail
		lastElmt.next = newElmt
		f.tail = newElmt
	}
	f.length++
}

/* */

/*
front
Returns the item at the front of the queue.
*/

func (f * Queue) front() (*Element) {
	return f.head
}

/*
dequeue()
Removes the item from the front of the queue.
*/
func (f * Queue) dequeue() (*Element) {
	if (f.length == 0){
		return nil
	} else {
		firstElmt := f.head
		f.head = f.head.next
		f.length--
		return firstElmt
	}
}

/*
isEmpty():Boolean
True if no more items can be dequeued and there is no front item.
*/

func (f *Queue) isEmpty() (bool) {
	if (f.head == nil){
		return true
	} else {
		return false
	}
}

func main(){
	l := createQueue()
	i := 0
	var e *Element
	for i != 10 {
		fmt.Println("Queued ",i*3)
		l.enqueue(&Element{body: i*3})
		i++
	}
	fmt.Println("-------------")

	for !l.isEmpty(){
		e = l.dequeue()
		fmt.Println("Popped ",e.body)
	}
}