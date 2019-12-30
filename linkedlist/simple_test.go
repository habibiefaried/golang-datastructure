package main 

import (
	"testing"
	"math/rand"
)

func TestSum(t *testing.T) {
    total := 10
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestCreate(t *testing.T){
	l := createList()
	if (l == nil){
		t.Errorf("Create failed!")
	}
}

func TestInsertFirst(t *testing.T){
	l := createList()
	i := 0
	var r int
	numEl := rand.Intn(500) // number of element

	for i < numEl {
		r = rand.Intn(numEl * 3)
		l.insertFirst(&Element{body: r})
		i++
	}

	if (l.length != numEl){
		t.Errorf("Some element missing!, only %d inserted",l.length)
	}
}

func BenchmarkInsertLast(b *testing.B){
	l := createList()
	i := 0
	var r int

	for i < b.N {
		r = rand.Intn(b.N * 3)
		l.insertLast(&Element{body: r})
		i++
	}
}